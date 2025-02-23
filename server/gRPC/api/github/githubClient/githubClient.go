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
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GithubClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	ActionsLauncher       *IServ.ActionLauncher
	cc                    gRPCService.GithubServiceClient
}

type WebHookRepoSendFunction = func(ctx context.Context, in *gRPCService.GitWebHookInfo, opts ...grpc.CallOption) (*gRPCService.GitWebHookInfo, error)

func NewGithubClient(conn *grpc.ClientConn) *GithubClient {
	micros := &IServ.ReactionLauncher{}
	launcher := &IServ.ActionLauncher{}
	git := &GithubClient{MicroservicesLauncher: micros, ActionsLauncher: launcher, cc: gRPCService.NewGithubServiceClient(conn)}
	(*git.MicroservicesLauncher)["updateRepo"] = git.updateRepository
	(*git.MicroservicesLauncher)["updateFile"] = git.updateFile
	(*git.MicroservicesLauncher)["deleteFile"] = git.deleteFile

	(*git.ActionsLauncher)["triggerPush"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreatePushWebhook)
	}
	(*git.ActionsLauncher)["triggerDelete"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateDeleteBranchWebhook)
	}
	(*git.ActionsLauncher)["triggerFork"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateForkRepositoryWebhook)
	}
	(*git.ActionsLauncher)["triggerCreate"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateNewBranchWebhook)
	}
	(*git.ActionsLauncher)["triggerIssue"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateNewIssueWebhook)
	}
	(*git.ActionsLauncher)["triggerIssueClose"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateDeleteIssueWebhook)
	}
	(*git.ActionsLauncher)["triggerPr"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateNewPRWebhook)
	}
	(*git.ActionsLauncher)["triggerPrClose"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateDeletePRWebhook)
	}
	(*git.ActionsLauncher)["triggerRelease"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateNewReleaseWebhook)
	}
	(*git.ActionsLauncher)["triggerReleaseDel"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return git.sendNewWebHookAction(scenario, actionId, userID, git.cc.CreateDeleteReleaseWebhook)
	}
	return git
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

func (git *GithubClient) updateRepository(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.Description, Datas: conv_utils.ConvertToMap[gRPCService.UpdateRepoInfos](&updateReq)}, nil
}

func (git *GithubClient) updateFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.Message, Datas: conv_utils.ConvertToMap[gRPCService.UpdateRepoFile](&updateReq)}, nil
}

func (git *GithubClient) deleteFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.Message, Datas: conv_utils.ConvertToMap[gRPCService.DeleteRepoFile](&updateReq)}, nil
}

func (git *GithubClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*git.ActionsLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such action microservice")
}

func (git *GithubClient) TriggerReaction(ingredients map[string]any, microservice string, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*git.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, userID)
	}
	return nil, errors.New("No such microservice")
}

func (git *GithubClient) TriggerWebhook(webhook *IServ.WebhookInfos, _ string, actionID int) (*IServ.WebHookResponseStatus, error) {
	payloadBytes, err := json.Marshal(webhook.Payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload send for this service")
	}
	if webhook.Header.Get("X-GitHub-Event") == "ping" {
		return &IServ.WebHookResponseStatus{Description: "ping"}, nil
	}
	headerBytes, err := json.Marshal(webhook.Header)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload send for this service")
	}
	if _, err := git.cc.TriggerWebHook(context.Background(), &gRPCService.GithubWebHookTriggerReq{
		ActionId: uint32(actionID),
		Payload:  payloadBytes,
		Header:   headerBytes}); err != nil {
		return nil, err
	}
	return &IServ.WebHookResponseStatus{Description: "Webhook triggered"}, nil
}

func (git *GithubClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := git.cc.SetActivateAction(ctx, &gRPCService.SetActivateGithub{
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

func (git *GithubClient) DeleteArea(ID uint, userID uint) (*IServ.DeleteResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(int(userID))
	_, err := git.cc.DeleteAction(ctx, &gRPCService.DeleteGithubActionReq{
		ActionId: uint32(ID),
	})
	if err != nil {
		return nil, err
	}
	return &IServ.DeleteResponseStatus{ID: ID}, nil
}
