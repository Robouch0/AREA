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
	"context"

	"google.golang.org/grpc"
)

type GitlabService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedGitlabServiceServer
}

func NewGitlabService() (*GitlabService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}

	return &GitlabService{tokenDb: tokenDb, reactService: nil}, err
}

func (git *GitlabService) InitReactClient(conn *grpc.ClientConn) {
	git.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (git *GitlabService) Test(ctx context.Context, req *gRPCService.TestReq) (*gRPCService.TestReq, error) {
	return nil, nil
}
