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
	"encoding/json"
	"log"

	"google.golang.org/grpc"
)

type ReactionServiceClient struct {
	gRPCService.ReactionServiceClient
}

func NewReactionServiceClient(conn *grpc.ClientConn) *ReactionServiceClient {
	return &ReactionServiceClient{gRPCService.NewReactionServiceClient(conn)}
}

func (react *ReactionServiceClient) SendAction(body []byte) (string, error) {
	// We want here to store the action from `body`
	react.LaunchReaction(context.Background(), &gRPCService.ReactionRequest{})
	return "", nil
}

////

type ReactionService struct {
	db      *db.UserDb
	clients map[string]ClientService

	gRPCService.UnimplementedReactionServiceServer
}

func NewReactionService(db *db.UserDb) ReactionService {
	return ReactionService{db: db, clients: make(map[string]ClientService)}
}

func (react *ReactionService) InitServiceClients(conn *grpc.ClientConn) {
	react.clients["dt"] = NewDateTimeServiceClient(conn)
	react.clients["hello"] = NewHelloServiceClient(conn)
}

func (react *ReactionService) LaunchReaction(_ context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	log.Println("Reaction searched")
	if service, ok := react.clients["hello"]; ok {
		log.Println("Reaction found and action sent")
		req := gRPCService.HelloWorldRequest{Message: req.Msg}
		b, err := json.Marshal(&req)
		if err != nil {
			log.Println("Error on jsonify")
			return nil, err
		}
		service.SendAction(b)
	}
	return &gRPCService.ReactionResponse{}, nil
}
