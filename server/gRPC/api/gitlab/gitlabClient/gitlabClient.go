//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabClient
//

package gitlab_client

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

type GitlabClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	ActionLauncher        *IServ.ActionLauncher

	cc gRPCService.GitlabServiceClient
}

func NewGitlabClient(conn *grpc.ClientConn) *GitlabClient {
	micros := &IServ.ReactionLauncher{}
	actions := &IServ.ActionLauncher{}
	gitlab := &GitlabClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewGitlabServiceClient(conn)}
	(*gitlab.MicroservicesLauncher)["createFile"] = gitlab.createFile
	(*gitlab.MicroservicesLauncher)["updateFile"] = gitlab.updateFile
	(*gitlab.MicroservicesLauncher)["deleteFile"] = gitlab.deleteFile
	(*gitlab.MicroservicesLauncher)["markItemDone"] = gitlab.markItemDone
	(*gitlab.MicroservicesLauncher)["markAllItemDone"] = gitlab.markAllItemDone
	return gitlab
}

func (git *GitlabClient) createFile(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.CreateLabRepoFileReq
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.CreateFile(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.CommitMessage}, nil
}

func (git *GitlabClient) updateFile(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.UpdateLabRepoFileReq
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.UpdateFile(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.CommitMessage}, nil
}

func (git *GitlabClient) deleteFile(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.DeleteLabRepoFileReq
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.DeleteFile(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.CommitMessage}, nil
}

func (git *GitlabClient) markItemDone(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.TodoLabItemDoneReq
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := git.cc.MarkItemAsDone(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Id}, nil
}

func (git *GitlabClient) markAllItemDone(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var updateReq gRPCService.AllTodoLabItemDoneReq
	err = json.Unmarshal(jsonString, &updateReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = git.cc.MarkAllItemAsDone(ctx, &updateReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Done"}, nil
}

func (git *GitlabClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*git.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such action microservice")
}

func (git *GitlabClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,
	prevOutput []byte,
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*git.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}

func (git *GitlabClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}

func (git *GitlabClient) DeactivateArea(microservice string, id uint, userID int) (*IServ.DeactivateResponseStatus, error) {
	return nil, status.Errorf(codes.Unavailable, "No Action Gitlab yet")
}
