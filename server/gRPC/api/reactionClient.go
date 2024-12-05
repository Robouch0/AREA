//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactionClient
//

package api

import (
	gRPCService "area/protogen/gRPC/proto"
	"context"

	"google.golang.org/grpc"
)

type ReactionServiceClient struct {
	gRPCService.ReactionServiceClient
}

func NewReactionServiceClient(conn *grpc.ClientConn) *ReactionServiceClient {
	return &ReactionServiceClient{gRPCService.NewReactionServiceClient(conn)}
}

func (react *ReactionServiceClient) SendAction(body []byte) (string, error) {
	// We want here to store the action from `body`
	react.LaunchReaction(context.Background(), &gRPCService.ReactionRequest{})
	return "", nil
}
