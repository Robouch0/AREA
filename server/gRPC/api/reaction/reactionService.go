//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionService
//

package reaction

import (
	"area/db"
	"area/gRPC/api/dateTime"
	"area/gRPC/api/hello"

	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

type ReactionService struct {
	db      *db.UserDb
	clients map[string]IServ.ClientService

	gRPCService.UnimplementedReactionServiceServer
}

func NewReactionService(db *db.UserDb) ReactionService {
	return ReactionService{db: db, clients: make(map[string]IServ.ClientService)}
}

func (react *ReactionService) InitServiceClients(conn *grpc.ClientConn) {
	react.clients["dt"] = dateTime.NewDateTimeServiceClient(conn)
	react.clients["hello"] = hello.NewHelloServiceClient(conn)
}

func (react *ReactionService) RegisterAction(_ context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	log.Println("Reaction searched")
	log.Println(req.Action, req.Reaction, req.UserId) // Do the database here
	// if _, ok := react.clients["hello"]; ok {
	// 	log.Println("Reaction found and action sent")
	// 	// req := gRPCService.HelloWorldRequest{Message: req.Msg}
	// 	// b := map[string]any{"msg": req.Message}
	// 	// _, err := service.SendAction(b) // TriggerReaction normally here
	// 	// if err != nil {
	// 	// log.Println(err)
	// 	// }
	// }
	return &gRPCService.ReactionResponse{Description: "Done", ActionId: 1}, nil
}
