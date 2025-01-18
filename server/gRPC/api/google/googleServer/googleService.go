//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleService
//

package google_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"

	"google.golang.org/grpc"
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
	_, err := grpcutils.GetUserIdFromContext(ctx, "google")
	if err != nil {
		return nil, err
	}
	// TODO Rahul:
	// Check if it is drive or gmail
	// Delete for GCP
	// Delete for the corresponding DB
	return req, nil
}
