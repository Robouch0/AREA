//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionClient
//

package reaction

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReactionServiceClient struct {
	gRPCService.ReactionServiceClient
}

func NewReactionServiceClient(conn *grpc.ClientConn) *ReactionServiceClient {
	return &ReactionServiceClient{gRPCService.NewReactionServiceClient(conn)}
}

// ReactionService is an internal service that is not exposed to public.
//
// Thus it has no "Public Status"
func (react *ReactionServiceClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	return nil, nil
}

func (react *ReactionServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No reaction available for this service")
}

func (_ *ReactionServiceClient) TriggerWebhook(webhook *IServ.WebhookInfos, _ string, _ int) (*IServ.WebHookResponseStatus, error) {
	return &IServ.WebHookResponseStatus{}, nil
}

func (react *ReactionServiceClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	bytesActIngredients, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	reactReq := &gRPCService.ReactionRequest{
		Action: &gRPCService.Action{
			Service:      scenario.Action.Service,
			Microservice: scenario.Action.Microservice,
			Ingredients:  bytesActIngredients,
		},
		Reactions: []*gRPCService.Reaction{},
	}
	for _, reaction := range scenario.Reactions {
		bytesReactIngredients, err := json.Marshal(reaction.Ingredients)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid ingredients given to the reaction")
		}
		reactReq.Reactions = append(reactReq.Reactions, &gRPCService.Reaction{
			Service:      reaction.Service,
			Microservice: reaction.Microservice,
			Ingredients:  bytesReactIngredients,
		})
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := react.RegisterAction(ctx, reactReq)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Description, ActionID: int(res.ActionId)}, nil
}

func (react *ReactionServiceClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := react.ReactionServiceClient.SetActivate(ctx, &gRPCService.AreaDeactivator{
		AreaId:    uint32(id),
		Activated: activated,
	})
	if err != nil {
		return nil, err
	}
	return &IServ.SetActivatedResponseStatus{
		ActionID:    id,
		Description: "DateTime Deactivated",
	}, nil
}
