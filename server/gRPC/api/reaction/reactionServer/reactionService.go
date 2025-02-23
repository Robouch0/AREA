//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionService
//

package reaction_server

import (
	"area/db"
	asana_client "area/gRPC/api/asana/asanaClient"
	crypto_client "area/gRPC/api/cryptoCompare/cryptoCompareClient"
	dateTime_client "area/gRPC/api/dateTime/dateTimeClient"
	discord_client "area/gRPC/api/discord/discordClient"
	github "area/gRPC/api/github/githubClient"
	gitlab_client "area/gRPC/api/gitlab/gitlabClient"
	google_client "area/gRPC/api/google/googleClient"
	huggingFace_client "area/gRPC/api/hugging_face/hugging_faceClient"
	miro_client "area/gRPC/api/miro/miroClient"
	IServ "area/gRPC/api/serviceInterface"
	spotify_client "area/gRPC/api/spotify/spotifyClient"
	weather_client "area/gRPC/api/weather/weatherClient"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	template_utils "area/utils/templateUtils"

	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	react.clients["dt"] = dateTime_client.NewDateTimeServiceClient(conn)
	react.clients["hf"] = huggingFace_client.NewHuggingFaceClient(conn)
	react.clients["github"] = github.NewGithubClient(conn)
	react.clients["gitlab"] = gitlab_client.NewGitlabClient(conn)
	react.clients["discord"] = discord_client.NewDiscordClient(conn)
	react.clients["google"] = google_client.NewGoogleClient(conn)
	react.clients["spotify"] = spotify_client.NewSpotifyClient(conn)
	react.clients["weather"] = weather_client.NewWeatherClient(conn)
	react.clients["asana"] = asana_client.NewAsanaClient(conn)
	react.clients["miro"] = miro_client.NewMiroClient(conn)
	react.clients["crypto"] = crypto_client.NewCryptoClient(conn)
}

func (react *ReactionService) LaunchReaction(ctx context.Context, req *gRPCService.LaunchRequest) (*gRPCService.LaunchResponse, error) {
	userID, errClaim := grpcutils.GetUserIdFromContext(ctx, "ReactionService")
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
	reactions, err := react.ReactionDB.GetReactionsByAreaID(area.ID) // Check if it is activated ?
	if reactions == nil {
		return nil, status.Errorf(codes.Internal, "no associated reaction")
	}

	var cache map[string]any
	errIngAct := json.Unmarshal(req.PrevOutput, &cache)
	if errIngAct != nil {
		return nil, err
	}

	for _, re := range *reactions {
		if cliService, ok := react.clients[re.Reaction.Service]; ok {
			template_utils.FormatIngredients(&re.Reaction.Ingredients, &cache)

			reactionRes, err := cliService.TriggerReaction(re.Reaction.Ingredients, re.Reaction.Microservice, int(area.UserID))
			if err != nil {
				fmt.Println("error: ", err)
				return nil, err
			}
			log.Println(reactionRes)

			utils.MergeMaps[any](&cache, &reactionRes.Datas)
		}
	}
	return &gRPCService.LaunchResponse{}, nil
}

func (react *ReactionService) RegisterAction(ctx context.Context, req *gRPCService.ReactionRequest) (*gRPCService.ReactionResponse, error) {
	userID, errClaim := grpcutils.GetUserIdFromContext(ctx, "ReactionService")
	if errClaim != nil {
		return nil, errClaim
	}

	newArea, err := react.AreaDB.InsertNewArea(userID, false)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var ingredientsAction map[string]any
	errIngAct := json.Unmarshal(req.Action.Ingredients, &ingredientsAction)
	if errIngAct != nil {
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

	for _, reaction := range req.Reactions {
		var ingredientsReaction map[string]any

		errIngReact := json.Unmarshal(reaction.Ingredients, &ingredientsReaction)
		if errIngReact != nil {
			log.Println(errIngReact)
			return nil, status.Errorf(codes.InvalidArgument, "Invalid ingredient given to reaction")
		}

		_, err := react.ReactionDB.InsertNewReaction(
			&models.Reaction{
				Service:      reaction.Service,
				Microservice: reaction.Microservice,
				Ingredients:  ingredientsReaction,
			},
			newArea.ID,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &gRPCService.ReactionResponse{Description: "Done", ActionId: int64(act.ID)}, nil
}

func (react *ReactionService) SetActivate(ctx context.Context, req *gRPCService.AreaDeactivator) (*gRPCService.AreaDeactivator, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "reaction")
	if err != nil {
		return nil, err
	}
	area, err := react.AreaDB.GetUserAreaByID(userID, uint(req.AreaId))
	if err != nil {
		return nil, err
	}
	log.Printf("Setting area (%v) activated to %v\n", req.AreaId, req.Activated)
	_, err = react.clients[area.Action.Action.Service].SetActivate(area.Action.Action.Microservice, area.Action.ID, int(userID), req.Activated)
	if err != nil {
		log.Println("Service Action error")
		return nil, err
	}
	_, err = react.AreaDB.SetActivateByAreaID(req.Activated, userID, uint(req.AreaId))
	if err != nil {
		log.Println("AreaDB error")
		return nil, err
	}
	return req, nil
}

func (react *ReactionService) DeleteUserArea(ctx context.Context, req *gRPCService.DeleteAreaReq) (*gRPCService.DeleteAreaReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "reaction")
	if err != nil {
		return nil, err
	}
	area, err := react.AreaDB.GetUserAreaByID(userID, uint(req.AreaId))
	if err != nil {
		return nil, err
	}

	_, errAction := react.clients[area.Action.Action.Service].DeleteArea(area.Action.ID, userID)
	errAct := react.ActionDB.DeleteByActionID(area.Action.ID)
	errReact := react.ReactionDB.DeleteByAreaID(area.ID)
	errArea := react.AreaDB.DeleteByID(userID, area.ID)
	if err := cmp.Or(errAction, errAct, errReact, errArea); err != nil {
		return nil, err
	}
	return req, nil
}
