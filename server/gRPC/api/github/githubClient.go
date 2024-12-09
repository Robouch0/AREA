//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubClient
//

package github

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
)

type GithubClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	cc                    gRPCService.GithubServiceClient
}

func NewGithubClient(conn *grpc.ClientConn) *GithubClient {
	micros := &IServ.MicroserviceLauncher{}
	git := &GithubClient{MicroservicesLauncher: micros, cc: gRPCService.NewGithubServiceClient(conn)}
	(*git.MicroservicesLauncher)["updateRepo"] = git.updateRepository
	return git
}

func (git *GithubClient) updateRepository(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.UpdateRepoReq
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	res, err := git.cc.UpdateRepository(context.Background(), &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Description}, nil
}

func (git *GithubClient) SendAction(body map[string]any, actionId int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in hugging face service (Next will be Webhooks)")
}

func (git *GithubClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*git.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput)
	}
	return nil, errors.New("No such microservice")
}
