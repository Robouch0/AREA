//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleService
//

package google

import (
	"area/db"
	"area/gRPC/api/google/gmail"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	sendMessageMeURL = "https://gmail.googleapis.com/gmail/v1/users/me/messages/send"
)

type GoogleService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGoogleServiceServer
}

func NewGoogleService() (*GoogleService, error) {
	tokenDb, err := db.InitTokenDb()

	return &GoogleService{tokenDb: tokenDb, reactService: nil}, err
}

func (google *GoogleService) SendEmailMe(ctx context.Context, req *gRPCService.EmailRequestMe) (*gRPCService.EmailRequestMe, error) {
	userID, errClaim := utils.GetUserIdFromContext(ctx, "GoogleGmailService")
	if errClaim != nil {
		return nil, errClaim
	}

	tokenInfo, err := google.tokenDb.GetUserTokenByProvider(int64(userID), "google")
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
