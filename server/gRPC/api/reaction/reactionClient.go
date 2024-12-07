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

	"google.golang.org/grpc"
)

type AreaScenario struct {
	UserId   int             `json:"user_id"`
	Action   models.Action   `json:"action"`
	Reaction models.Reaction `json:"reaction"`
}

type ReactionServiceClient struct {
	gRPCService.ReactionServiceClient
}

func NewReactionServiceClient(conn *grpc.ClientConn) *ReactionServiceClient {
	return &ReactionServiceClient{gRPCService.NewReactionServiceClient(conn)}
}

func (react *ReactionServiceClient) SendAction(body map[string]any, actionID int) (*IServ.ActionResponseStatus, error) {
	jsonString, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var scenarioArea AreaScenario
	err = json.Unmarshal(jsonString, &scenarioArea)
	if err != nil {
		return nil, err
	}

	bytesActIngredients, err := json.Marshal(scenarioArea.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	bytesReactIngredients, err := json.Marshal(scenarioArea.Reaction.Ingredients)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	res, err := react.RegisterAction(ctx, &gRPCService.ReactionRequest{
		UserId: int64(scenarioArea.UserId),
		Action: &gRPCService.Action{
			Service:      scenarioArea.Action.Service,
			Microservice: scenarioArea.Action.Microservice,
			Ingredients:  bytesActIngredients,
		},
		Reaction: &gRPCService.Reaction{
			Service:      scenarioArea.Reaction.Service,
			Microservice: scenarioArea.Reaction.Microservice,
			Ingredients:  bytesReactIngredients,
		}})
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Description, ActionID: int(res.ActionId)}, nil
}
