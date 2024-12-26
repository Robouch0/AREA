//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleService
//

package google_server

import (
	"area/db"
	"area/gRPC/api/google/gmail"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoogleService struct {
	tokenDb      *db.TokenDb
	gmailDb      *db.GoogleGmailDB
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGoogleServiceServer
}

func NewGoogleService() (*GoogleService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}

	gmailDB, err := db.InitGoogleGmailDb()
	if err != nil {
		return nil, err
	}
	return &GoogleService{tokenDb: tokenDb, gmailDb: gmailDB, reactService: nil}, err
}

func (google *GoogleService) InitReactClient(conn *grpc.ClientConn) {
	google.reactService = gRPCService.NewReactionServiceClient(conn)
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
		HistoryID:   watchMeResponse.HistoryID,
		EmailAdress: gTokenInfo.Email,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot store new data in DB: %v", err)
	}
	return req, nil
}

func (google *GoogleService) WatchMeTrigger(ctx context.Context, req *gRPCService.GmailTriggerReq) (*gRPCService.GmailTriggerReq, error) {
	var payload gmail.PubSubPayload
	if json.Unmarshal(req.Payload, &payload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}
	data, err := utils.DecodeBase64ToStruct[gmail.GmailPayload]([]byte(payload.Message.Data))
	if err != nil {
		log.Println("Cannot convert to struct")
		return nil, err
	}
	act, err := google.gmailDb.GetByEmail(data.EmailAddress)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if act.Activated {
		ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
		_, err := google.reactService.LaunchReaction(
			ctx,
			&gRPCService.LaunchRequest{ActionId: int64(act.ActionID), PrevOutput: []byte(payload.Message.Data)},
		)
		if err != nil {
			return nil, err
		}
	}
	return req, nil
}
