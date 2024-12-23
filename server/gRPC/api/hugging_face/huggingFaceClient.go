//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceClient
//

package huggingFace

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
)

type HuggingFaceServiceClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	ActionsLauncher       *IServ.ActionLauncher

	cc gRPCService.HuggingFaceServiceClient
}

// Alias for the webhook send function for HuggingFace
type WebHookRepoSendFunction = func(
	ctx context.Context, in *gRPCService.HFWebHookInfo, opts ...grpc.CallOption) (*gRPCService.HFWebHookInfo, error)

func NewHuggingFaceClient(conn *grpc.ClientConn) *HuggingFaceServiceClient {
	hfCli := &HuggingFaceServiceClient{
		MicroservicesLauncher: &IServ.MicroserviceLauncher{},
		ActionsLauncher:       &IServ.ActionLauncher{},
		cc:                    gRPCService.NewHuggingFaceServiceClient(conn),
	}
	(*hfCli.MicroservicesLauncher)["textGen"] = hfCli.sendTextGenerationReaction

	(*hfCli.ActionsLauncher)["push"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return hfCli.sendNewWebHookAction(scenario, actionId, userID, hfCli.cc.CreateRepoUpdateWebHook)
	}
	(*hfCli.ActionsLauncher)["pr"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return hfCli.sendNewWebHookAction(scenario, actionId, userID, hfCli.cc.CreateNewPRWebHook)
	}
	(*hfCli.ActionsLauncher)["discussion"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return hfCli.sendNewWebHookAction(scenario, actionId, userID, hfCli.cc.CreateNewDiscussionWebHook)
	}
	return hfCli
}

func (hfCli *HuggingFaceServiceClient) sendNewWebHookAction(
	scenario models.AreaScenario, actionID, userID int, sendFn WebHookRepoSendFunction,
) (*IServ.ActionResponseStatus, error) {
	webHookIng, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	webHookReq := gRPCService.HFWebHookInfo{ActionId: int32(actionID)}
	err = json.Unmarshal(webHookIng, &webHookReq)
	if err != nil {
		return nil, err
	}
	ctx := utils.CreateContextFromUserID(userID)
	res, err := sendFn(ctx, &webHookReq)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Name}, nil
}

func (hfCli *HuggingFaceServiceClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*hfCli.ActionsLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such action microservice")
}

func (hfCli *HuggingFaceServiceClient) sendTextGenerationReaction(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	ctx := utils.CreateContextFromUserID(userID)
	res, err := hfCli.cc.LaunchTextGeneration(ctx, &gRPCService.TextGenerationReq{Inputs: ingredients["inputs"].(string)})
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.GeneratedText}, nil
}

func (hfCli *HuggingFaceServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*hfCli.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}
