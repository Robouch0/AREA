//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleGmailEmail
//

package google_server

import (
	"area/gRPC/api/google/gmail"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"bytes"
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

	emailRawContent := gmail.CreateEmailRawContent("dikrah25@gmail.com", req.To, req.Subject, req.BodyMessage)
	emailBody := gmail.NewEmailRequestBody(emailRawContent)
	b, err := json.Marshal(&emailBody)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to convert the content to bytes"))
	}

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
	return nil, nil
}
