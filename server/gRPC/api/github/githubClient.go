//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubClient
//

package github

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GithubClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	ActionsLauncher       *IServ.ActionLauncher
	cc                    gRPCService.GithubServiceClient
}

type WebHookRepoSendFunction = func(ctx context.Context, in *gRPCService.GitWebHookInfo, opts ...grpc.CallOption) (*gRPCService.GitWebHookInfo, error)

func NewGithubClient(conn *grpc.ClientConn) *GithubClient {
	micros := &IServ.MicroserviceLauncher{}
	launcher := &IServ.ActionLauncher{}
	git := &GithubClient{MicroservicesLauncher: micros, ActionsLauncher: launcher, cc: gRPCService.NewGithubServiceClient(conn)}
	(*git.MicroservicesLauncher)["updateRepo"] = git.updateRepository
	(*git.MicroservicesLauncher)["updateFile"] = git.updateFile
	(*git.MicroservicesLauncher)["deleteFile"] = git.deleteFile

	(*git.ActionsLauncher)["triggerPush"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreatePushWebhook)
	}
	return git
}

func (git *GithubClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Github",
		RefName: "github",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Update Repository Informations",
				RefName: "updateRepo",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"name": {
						Value:       "",
						Type:        "string",
						Description: "New name for the repository",
						Required:    true,
					},
					"description": {
						Value:       "",
						Type:        "string",
						Description: "New description for the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update a file in a repository",
				RefName: "updateFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"path": {
						Value:       "",
						Type:        "string",
						Description: "Path of the file in the repository",
						Required:    true,
					},
					"message": {
						Value:       "",
						Type:        "string",
						Description: "Commit message",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "New content of the file",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete Repository File",
				RefName: "deleteFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"path": {
						Value:       "",
						Type:        "string",
						Description: "Path of the file in the repository",
						Required:    true,
					},
					"message": {
						Value:       "",
						Type:        "string",
						Description: "Commit message",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}

func (git *GithubClient) sendNewWebHookAction(
	scenario models.AreaScenario, actionID, userID int, sendFn WebHookRepoSendFunction,
) (*IServ.ActionResponseStatus, error) {
	webHookIng, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	webHookReq := gRPCService.GitWebHookInfo{ActionId: int32(actionID)}
	err = json.Unmarshal(webHookIng, &webHookReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := sendFn(ctx, &webHookReq)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Repo}, nil
}

func (git *GithubClient) updateRepository(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.UpdateRepoInfos
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.UpdateRepository(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Description}, nil
}

func (git *GithubClient) updateFile(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.UpdateRepoFile
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.UpdateFile(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Message}, nil
}

func (git *GithubClient) deleteFile(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteRepoFile
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.DeleteFile(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Message}, nil
}

func (git *GithubClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*git.ActionsLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such action microservice")
}

func (git *GithubClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*git.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}

func (git *GithubClient) TriggerWebhook(payload map[string]any, _ string, actionID int) (*IServ.WebHookResponseStatus, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload send for this service")
	}
	if _, err := git.cc.TriggerWebHook(context.Background(), &gRPCService.GithubWebHookTriggerReq{ActionId: uint32(actionID), Payload: payloadBytes}); err != nil {
		return nil, err
	}
	return &IServ.WebHookResponseStatus{Description: "Webhook triggered"}, nil
}
