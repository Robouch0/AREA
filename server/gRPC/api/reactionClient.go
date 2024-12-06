//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionClient
//

package api

import (
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

func (react *ReactionServiceClient) SendAction(body map[string]any) (string, error) {
	jsonString, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	var scenarioArea AreaScenario
	err = json.Unmarshal(jsonString, &scenarioArea)
	if err != nil {
		return "", err
	}

	bytesActIngredients, err := json.Marshal(scenarioArea.Action.Ingredients)
	if err != nil {
		return "", err
	}

	bytesReactIngredients, err := json.Marshal(scenarioArea.Reaction.Ingredients)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	react.RegisterAction(ctx, &gRPCService.ReactionRequest{
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
	return "", nil
}
