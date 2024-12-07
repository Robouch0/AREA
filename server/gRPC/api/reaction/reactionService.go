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
	"area/models"
	"cmp"
	"encoding/json"

	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

type ReactionService struct {
	AreaDB     *db.AreaDB
	ActionDB   *db.ActionsDb
	ReactionDB *db.ReactionDb
	clients    map[string]IServ.ClientService

	gRPCService.UnimplementedReactionServiceServer
}

func NewReactionService() (*ReactionService, error) {
	AreaDB, errAreaDB := db.InitAreaDb()
	ActionDB, errActionDB := db.InitActionsDb()
	ReactionDB, errReactionDB := db.InitReactionsDb()

	if err := cmp.Or(errAreaDB, errActionDB, errReactionDB); err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	return &ReactionService{
		AreaDB:     AreaDB,
		ActionDB:   ActionDB,
		ReactionDB: ReactionDB,
		clients:    make(map[string]IServ.ClientService),
	}, nil
}

func (react *ReactionService) InitServiceClients(conn *grpc.ClientConn) {
	react.clients["dt"] = dateTime.NewDateTimeServiceClient(conn)
	react.clients["hello"] = hello.NewHelloServiceClient(conn)
}

func (react *ReactionService) RegisterAction(_ context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	log.Println("Reaction searched")

	newArea, err := react.AreaDB.InsertNewArea(uint(req.UserId), false)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var ingredientsAction map[string]any
	var ingredientsReaction map[string]any
	errIngAct := json.Unmarshal(req.Action.Ingredients, &ingredientsAction) // Check error here !
	errIngReact := json.Unmarshal(req.Reaction.Ingredients, &ingredientsReaction)
	if err = cmp.Or(errIngAct, errIngReact); err != nil {
		log.Println("Here\n", err)
		return nil, err
	}
	log.Println(ingredientsAction)
	act, err := react.ActionDB.InsertNewAction(
		&models.Action{
			Service:      req.Action.Service,
			Microservice: req.Action.Microservice,
			Ingredients:  ingredientsAction,
		},
		newArea.ID,
	)
	if err != nil {
		log.Println("Here new Action: ", err)
		return nil, err
	}
	log.Printf("This is the action: %v\n", act) // It works normally
	re, err := react.ReactionDB.InsertNewReaction(
		&models.Reaction{
			Service:      req.Reaction.Service,
			Microservice: req.Reaction.Microservice,
			Ingredients:  ingredientsReaction,
		},
		newArea.ID,
	)
	if err != nil {
		log.Println("Here new Reaction: ", err)
		return nil, err
	}
	log.Printf("This is the reaction: %v\n", re)

	log.Println(react.AreaDB.GetAreaByID(newArea.ID))
	return &gRPCService.ReactionResponse{Description: "Done", ActionId: 1}, nil
}
