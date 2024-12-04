//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionService
//

package api

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"context"

	"google.golang.org/grpc"
)

type ReactionServiceClient struct {
	gRPCService.ReactionServiceClient
}

func NewReactionServiceClient(conn *grpc.ClientConn) *ReactionServiceClient {
	return &ReactionServiceClient{gRPCService.NewReactionServiceClient(conn)}
}

func (react *ReactionServiceClient) SendAction(body []byte) (string, error) {
	// msg := new(msg)
	// err := json.Unmarshal([]byte(body), msg)
	//
	// r, err := hello.SayHello(context.Background(), &gRPCService.HelloWorldRequest{Message: msg.Msg})
	// if err != nil {
	// return "", err
	// }
	// return r.GetMessage(), nil
	react.LaunchReaction(context.Background(), &gRPCService.ReactionRequest{})
	return "", nil
}

////

type ReactionService struct {
	db      *db.UserDb
	clients map[string]ClientService
	gRPCService.UnimplementedDateTimeServiceServer
}

func NewReactionService(db *db.UserDb) ReactionService {
	return ReactionService{db: db, clients: make(map[string]ClientService)}
}

func (react *ReactionService) LaunchReaction(_ context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	return &gRPCService.ReactionResponse{}, nil
}
