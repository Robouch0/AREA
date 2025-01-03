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

	bytesReactIngredients, err := json.Marshal(scenario.Reaction.Ingredients)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := react.RegisterAction(
		ctx,
		&gRPCService.ReactionRequest{
			Action: &gRPCService.Action{
				Service:      scenario.Action.Service,
				Microservice: scenario.Action.Microservice,
				Ingredients:  bytesActIngredients,
			},
			Reaction: &gRPCService.Reaction{
				Service:      scenario.Reaction.Service,
				Microservice: scenario.Reaction.Microservice,
				Ingredients:  bytesReactIngredients,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Description, ActionID: int(res.ActionId)}, nil
}

func (react *ReactionServiceClient) DeactivateArea(microservice string, id uint, userID int) (*IServ.DeactivateResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := react.ReactionServiceClient.DeactivateArea(ctx, &gRPCService.AreaDeactivator{
		AreaId: uint32(id),
	})
	if err != nil {
		return nil, err
	}
	return &IServ.DeactivateResponseStatus{
		ActionID:    id,
		Description: "DateTime Deactivated",
	}, nil
}
