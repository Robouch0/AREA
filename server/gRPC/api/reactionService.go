//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionService
//

package api

import (
	"area/db"
	"area/gRPC/api/dateTime"
	"area/gRPC/api/hello"

	gRPCService "area/protogen/gRPC/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

type ReactionService struct {
	db      *db.UserDb
	clients map[string]ClientService

	gRPCService.UnimplementedReactionServiceServer
}

func NewReactionService(db *db.UserDb) ReactionService {
	return ReactionService{db: db, clients: make(map[string]ClientService)}
}

func (react *ReactionService) InitServiceClients(conn *grpc.ClientConn) {
	react.clients["dt"] = dateTime.NewDateTimeServiceClient(conn)
	react.clients["hello"] = hello.NewHelloServiceClient(conn)
}

func (react *ReactionService) LaunchReaction(_ context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	log.Println("Reaction searched")
	if service, ok := react.clients["hello"]; ok {
		log.Println("Reaction found and action sent")
		req := gRPCService.HelloWorldRequest{Message: req.Msg}
		b := map[string]any{"msg": req.Message}
		_, err := service.SendAction(b) // TriggerReaction normally here
		if err != nil {
			log.Println(err)
		}
	}
	return &gRPCService.ReactionResponse{}, nil
}
