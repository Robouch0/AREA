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
	"area/gRPC/api/discord"
	"area/gRPC/api/github"
	"area/gRPC/api/google"
	"area/gRPC/api/hello"
	"area/gRPC/api/spotify"
	huggingFace "area/gRPC/api/hugging_face"
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"

	"cmp"
	"context"
	"encoding/json"
	"fmt"
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
	react.clients["hf"] = huggingFace.NewHuggingFaceClient(conn)
	react.clients["github"] = github.NewGithubClient(conn)
	react.clients["discord"] = discord.NewDiscordClient(conn)
	react.clients["google"] = google.NewGoogleClient(conn)
	react.clients["spotify"] = spotify.NewSpotifyClient(conn)
}

func (react *ReactionService) LaunchReaction(ctx context.Context, req *gRPCService.LaunchRequest) (*gRPCService.LaunchResponse, error) {
	userID, errClaim := utils.GetUserIdFromContext(ctx, "ReactionService")
	if errClaim != nil {
		return nil, errClaim
	}

	log.Println("ActionID Launch reaction (reactService): ", req.ActionId)
	area, err := react.AreaDB.GetUserAreaByActionID(userID, uint(req.ActionId))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("AreaID Launch: ", area.ID)
	reactions, err := react.ReactionDB.GetReactionsByAreaID(area.ID)
	for _, re := range *reactions {
		if cliService, ok := react.clients[re.Reaction.Service]; ok {
			res, err := cliService.TriggerReaction(re.Reaction.Ingredients, re.Reaction.Microservice, req.PrevOutput, int(area.UserID))
			if err != nil {
				fmt.Println("error: ", err)
				return nil, err
			}
			log.Println(res)
		}
	}
	return &gRPCService.LaunchResponse{}, nil
}

func (react *ReactionService) RegisterAction(ctx context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	userID, errClaim := utils.GetUserIdFromContext(ctx, "ReactionService")
	if errClaim != nil {
		return nil, errClaim
	}

	newArea, err := react.AreaDB.InsertNewArea(userID, false)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var ingredientsAction map[string]any
	var ingredientsReaction map[string]any
	errIngAct := json.Unmarshal(req.Action.Ingredients, &ingredientsAction)
	errIngReact := json.Unmarshal(req.Reaction.Ingredients, &ingredientsReaction)
	if err = cmp.Or(errIngAct, errIngReact); err != nil {
		return nil, err
	}

	act, err := react.ActionDB.InsertNewAction(
		&models.Action{
			Service:      req.Action.Service,
			Microservice: req.Action.Microservice,
			Ingredients:  ingredientsAction,
		},
		newArea.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = react.ReactionDB.InsertNewReaction(
		&models.Reaction{
			Service:      req.Reaction.Service,
			Microservice: req.Reaction.Microservice,
			Ingredients:  ingredientsReaction,
		},
		newArea.ID,
	)
	if err != nil {
		return nil, err
	}
	return &gRPCService.ReactionResponse{Description: "Done", ActionId: int64(act.ID)}, nil
}
