//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleService
//

package google_server

import (
	"area/db"
	"area/gRPC/api/google/drive"
	"area/gRPC/api/google/gmail"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoogleService struct {
	tokenDb      *db.TokenDb
	gmailDb      *db.GoogleGmailDB
	driveDb      *db.GoogleDriveDB
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

	driveDB, err := db.InitGoogleDriveDB()
	if err != nil {
		return nil, err
	}
	return &GoogleService{tokenDb: tokenDb, gmailDb: gmailDB, driveDb: driveDB, reactService: nil}, err
}

func (google *GoogleService) InitReactClient(conn *grpc.ClientConn) {
	google.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (google *GoogleService) DeleteAction(ctx context.Context, req *gRPCService.DeleteGoogleActionReq) (*gRPCService.DeleteGoogleActionReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, google.tokenDb, "GoogleGmailService", "google")
	if err != nil {
		return nil, err
	}

	if driveData, err := google.driveDb.GetActionByID(strconv.Itoa(int(req.ActionId))); err == nil {
		drive.StopWatchDrive(tokenInfo, driveData.ChannelID, driveData.ResourceID)
		google.driveDb.DeleteByActionID(uint(tokenInfo.UserID), uint(req.ActionId))
	} else if _, err := google.gmailDb.GetActionByID(uint(req.ActionId)); err == nil {
		gmail.StopPubSub(tokenInfo)
		google.driveDb.DeleteByActionID(uint(tokenInfo.UserID), uint(req.ActionId))
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid actionID sent")
	}
	return req, nil
}
