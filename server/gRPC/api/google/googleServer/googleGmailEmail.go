//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleGmailEmail
//

package google_server

import (
	"area/gRPC/api/google/gmail"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	sendMessageMeURL = "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"
)

func (google *GoogleService) DeleteEmailMe(ctx context.Context, req *gRPCService.DeleteEmailRequestMe) (*gRPCService.DeleteEmailRequestMe, error) {
	if req.Subject == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Empty subject for email deletion")
	}
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleGmailService", "google")
	if err != nil {
		return nil, err
	}

	emails, err := gmail.GetListEmails("me", tokenInfo.AccessToken, "")
	if err != nil {
		return nil, err
	}

	for _, message := range emails.Messages {
		mess, err := gmail.GetEmail("me", tokenInfo.AccessToken, message.ID, "metadata", "subject")
		if err != nil {
			return nil, nil
		}
		idx := slices.IndexFunc[[]gmail.GmailHeader](mess.Payload.Headers, func(h gmail.GmailHeader) bool {
			return h.Name == "Subject"
		})
		if idx == -1 {
			continue
		}
		if mess.Payload.Headers[idx].Value == req.Subject {
			err := gmail.DeleteEmail("me", tokenInfo.AccessToken, message.ID)
			if err != nil {
				log.Println("Delete Error: ", err)
				return nil, err
			}
			return req, nil
		}
	}
	return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Did not find any email with subject %s", req.Subject))
}

func (google *GoogleService) SendEmailMe(ctx context.Context, req *gRPCService.EmailRequestMe) (*gRPCService.EmailRequestMe, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleGmailService", "google")
	if err != nil {
		return nil, err
	}
	emailMeInfo, err := GetTokenInfo(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}

	emailRawContent := gmail.CreateEmailRawContent(emailMeInfo.Email, req.To, req.Subject, req.BodyMessage)
	emailBody := gmail.NewEmailRequestBody(emailRawContent)
	b, err := json.Marshal(&emailBody)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

	log.Println("Email sent to: ", req.To)
	postRequest, err := http.NewRequest("POST", sendMessageMeURL, bytes.NewBuffer(b))
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	postRequest.Header.Add("Content-Type", "message/rfc822")
	postRequest.Header.Add("Accept", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		io.Copy(os.Stderr, resp.Body)
		return nil, status.Errorf(codes.Aborted, resp.Status)
	}
	log.Println(resp.Body)
	return req, nil
}

func (google *GoogleService) WatchGmailEmail(ctx context.Context, req *gRPCService.EmailTriggerReq) (*gRPCService.EmailTriggerReq, error) {
	tokenInfo, errClaim := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleService", "google")
	if errClaim != nil {
		return nil, errClaim
	}
	watchMeResponse, err := gmail.SendWatchMeRequest(tokenInfo)
	if err != nil {
		return nil, err
	}
	gTokenInfo, err := GetTokenInfo(tokenInfo.AccessToken)
	if err != nil {
		return nil, err
	}

	_, err = google.gmailDb.StoreNewGWatch(&models.Gmail{
		ActionID:    uint(req.ActionId),
		UserID:      uint(tokenInfo.UserID),
		Activated:   true,
		FirstTime:   true,
		HistoryID:   watchMeResponse.HistoryID,
		EmailAdress: gTokenInfo.Email,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot store new data in DB: %v", err)
	}
	log.Println("New Watch for gmail")
	return req, nil
}

func (google *GoogleService) WatchMeTrigger(ctx context.Context, req *gRPCService.GmailTriggerReq) (*gRPCService.GmailTriggerReq, error) {
	var payload gmail.PubSubPayload
	if json.Unmarshal(req.Payload, &payload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	data, err := utils.DecodeBase64ToStruct[gmail.GmailPayload]([]byte(payload.Message.Data))
	if err != nil {
		log.Println("Cannot convert gmail payload to struct")
		return nil, err
	}
	act, err := google.gmailDb.GetByEmail(data.EmailAddress)
	if err != nil {
		log.Println("No such adress mail: ", data.EmailAddress)
		return nil, err
	}
	if act.Activated {
		if act.FirstTime {
			log.Println("First Time request for gmail (On Watch activation)")
			_, err := google.gmailDb.SetFirstTime(act.ActionID, false)
			return req, err
		}

		bytesMess, err := json.Marshal(payload.Message)
		if err != nil {
			return nil, err
		}

		log.Println("Action (GmailWatchMe) is activated so redirect to reaction Service")
		ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err = google.reactService.LaunchReaction(
			ctx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: bytesMess},
		)
		if err != nil {
			log.Println("Error ReactionService", err)
			return nil, err
		}
	}
	log.Println("Watchme trigger !")
	return req, nil
}

func (google *GoogleService) SetActivateGmailAction(ctx context.Context, req *gRPCService.SetActivateGmail) (*gRPCService.SetActivateGmail, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleGmailService", "google")
	if err != nil {
		return nil, err
	}
	if !req.Activated {
		err = gmail.StopPubSub(tokenInfo)
	} else {
		_, errWatch := gmail.SendWatchMeRequest(tokenInfo)
		_, errFt := google.gmailDb.SetFirstTime(uint(req.ActionId), true)
		err = cmp.Or(errFt, errWatch)
	}
	if err != nil {
		return nil, err
	}
	_, err = google.gmailDb.SetActivateByActionID(req.Activated, uint(tokenInfo.UserID), uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}
