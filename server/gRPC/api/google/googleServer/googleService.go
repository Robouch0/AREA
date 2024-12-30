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

	"google.golang.org/grpc"
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
