//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabServer
//

package gitlab_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"

	"google.golang.org/grpc"
)

type GoogleService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGitlabServiceServer
}

func NewGoogleService() (*GoogleService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}

	return &GoogleService{tokenDb: tokenDb, reactService: nil}, err
}

func (git *GoogleService) InitReactClient(conn *grpc.ClientConn) {
	git.reactService = gRPCService.NewReactionServiceClient(conn)
}
