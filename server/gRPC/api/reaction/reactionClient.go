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
	"context"
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

func (react *ReactionServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No reaction available for this service")
}

func (react *ReactionServiceClient) SendAction(scenario models.AreaScenario, actionID int) (*IServ.ActionResponseStatus, error) {
	bytesActIngredients, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	bytesReactIngredients, err := json.Marshal(scenario.Reaction.Ingredients)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	res, err := react.RegisterAction(ctx, &gRPCService.ReactionRequest{
		UserId: int64(scenario.UserId),
		Action: &gRPCService.Action{
			Service:      scenario.Action.Service,
			Microservice: scenario.Action.Microservice,
			Ingredients:  bytesActIngredients,
		},
		Reaction: &gRPCService.Reaction{
			Service:      scenario.Reaction.Service,
			Microservice: scenario.Reaction.Microservice,
			Ingredients:  bytesReactIngredients,
		}})
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Description, ActionID: int(res.ActionId)}, nil
}
