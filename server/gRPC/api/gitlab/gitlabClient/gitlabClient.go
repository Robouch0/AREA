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
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	"context"
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

type WebHookRepoSendFunction = func(ctx context.Context, in *gRPCService.GitlabWebHookInfo, opts ...grpc.CallOption) (*gRPCService.GitlabWebHookInfo, error)

func NewGitlabClient(conn *grpc.ClientConn) *GitlabClient {
	micros := &IServ.ReactionLauncher{}
	actions := &IServ.ActionLauncher{}
	gitlab := &GitlabClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewGitlabServiceClient(conn)}
	(*gitlab.MicroservicesLauncher)["createFile"] = gitlab.createFile
	(*gitlab.MicroservicesLauncher)["updateFile"] = gitlab.updateFile
	(*gitlab.MicroservicesLauncher)["deleteFile"] = gitlab.deleteFile
	(*gitlab.MicroservicesLauncher)["markItemDone"] = gitlab.markItemDone
	(*gitlab.MicroservicesLauncher)["markAllItemDone"] = gitlab.markAllItemDone

	(*gitlab.ActionLauncher)["triggerPush"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return gitlab.sendNewWebHookAction(scenario, actionId, userID, gitlab.cc.CreatePushWebhook)
	}
	(*gitlab.ActionLauncher)["triggerIssue"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return gitlab.sendNewWebHookAction(scenario, actionId, userID, gitlab.cc.CreateIssueWebhook)
	}
	(*gitlab.ActionLauncher)["triggerTag"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return gitlab.sendNewWebHookAction(scenario, actionId, userID, gitlab.cc.CreateTagWebhook)
	}
	(*gitlab.ActionLauncher)["triggerRelease"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return gitlab.sendNewWebHookAction(scenario, actionId, userID, gitlab.cc.CreateReleaseWebhook)
	}
	(*gitlab.ActionLauncher)["triggerMerge"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return gitlab.sendNewWebHookAction(scenario, actionId, userID, gitlab.cc.CreateMergeEventWebhook)
	}
	return gitlab
}

func (git *GitlabClient) sendNewWebHookAction(
	scenario models.AreaScenario, actionID, userID int, sendFn WebHookRepoSendFunction,
) (*IServ.ActionResponseStatus, error) {
	webHookIng, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	webHookReq := gRPCService.GitlabWebHookInfo{ActionId: int32(actionID)}
	err = json.Unmarshal(webHookIng, &webHookReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := sendFn(ctx, &webHookReq)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Id}, nil
}

func (git *GitlabClient) createFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.CommitMessage, Datas: conv_utils.ConvertToMap[gRPCService.CreateLabRepoFileReq](&updateReq)}, nil
}

func (git *GitlabClient) updateFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.CommitMessage, Datas: conv_utils.ConvertToMap[gRPCService.UpdateLabRepoFileReq](&updateReq)}, nil
}

func (git *GitlabClient) deleteFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.CommitMessage, Datas: conv_utils.ConvertToMap[gRPCService.DeleteLabRepoFileReq](&updateReq)}, nil
}

func (git *GitlabClient) markItemDone(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.Id, Datas: conv_utils.ConvertToMap[gRPCService.TodoLabItemDoneReq](&updateReq)}, nil
}

func (git *GitlabClient) markAllItemDone(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: "Done", Datas: conv_utils.ConvertToMap[gRPCService.AllTodoLabItemDoneReq](&updateReq)}, nil
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
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*git.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, userID)
	}
	return nil, errors.New("No such microservice")
}

func (git *GitlabClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	payloadBytes, err := json.Marshal(webhook.Payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload send for this service")
	}
	if _, err := git.cc.TriggerWebHook(context.Background(), &gRPCService.GitlabWebHookTriggerReq{ActionId: uint32(actionID), Payload: payloadBytes}); err != nil {
		return nil, err
	}
	return &IServ.WebHookResponseStatus{Description: "Webhook triggered"}, nil
}

func (git *GitlabClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := git.cc.SetActivateAction(ctx, &gRPCService.SetActivateGitlab{
		ActionId:  uint32(id),
		Activated: activated,
	})
	if err != nil {
		return nil, err
	}
	return &IServ.SetActivatedResponseStatus{
		ActionID:    id,
		Description: "Github Deactivated",
	}, nil
}

func (git *GitlabClient) DeleteArea(ID uint, userID uint) (*IServ.DeleteResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(int(userID))
	_, err := git.cc.DeleteAction(ctx, &gRPCService.DeleteGitlabActionReq{
		ActionId: uint32(ID),
	})
	if err != nil {
		return nil, err
	}
	return &IServ.DeleteResponseStatus{ID: ID}, nil
}
