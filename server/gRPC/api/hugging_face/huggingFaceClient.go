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
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HuggingFaceServiceClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	ActionsLauncher       *IServ.ActionLauncher

	cc gRPCService.HuggingFaceServiceClient
}

// Alias for the webhook send function for HuggingFace
type WebHookRepoSendFunction = func(ctx context.Context, in *gRPCService.HFWebHookInfo, opts ...grpc.CallOption) (*gRPCService.HFWebHookInfo, error)

func NewHuggingFaceClient(conn *grpc.ClientConn) *HuggingFaceServiceClient {
	hf := &HuggingFaceServiceClient{
		MicroservicesLauncher: &IServ.MicroserviceLauncher{},
		ActionsLauncher:       &IServ.ActionLauncher{},
		cc:                    gRPCService.NewHuggingFaceServiceClient(conn),
	}
	(*hf.MicroservicesLauncher)["textGen"] = hf.sendTextGenerationReaction

	(*hf.ActionsLauncher)["push"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return hf.sendNewWebHookAction(scenario, actionId, userID, hf.cc.CreateRepoUpdateWebHook)
	}
	(*hf.ActionsLauncher)["pr"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return hf.sendNewWebHookAction(scenario, actionId, userID, hf.cc.CreateNewPRWebHook)
	}
	(*hf.ActionsLauncher)["discussion"] = func(scenario models.AreaScenario, actionId, userID int) (*IServ.ActionResponseStatus, error) {
		return hf.sendNewWebHookAction(scenario, actionId, userID, hf.cc.CreateNewDiscussionWebHook)
	}
	return hf
}

func (hf *HuggingFaceServiceClient) sendNewWebHookAction(
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
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := sendFn(ctx, &webHookReq)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{Description: res.Name}, nil
}

func (hf *HuggingFaceServiceClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*hf.ActionsLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such action microservice")
}

func (hf *HuggingFaceServiceClient) sendTextGenerationReaction(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.LaunchTextGeneration(ctx, &gRPCService.TextGenerationReq{Inputs: ingredients["inputs"].(string)})
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.GeneratedText}, nil
}

func (hf *HuggingFaceServiceClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*hf.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	return nil, errors.New("No such microservice")
}

// No need to check microservice here but later there will be a map for multiple type of actions (like actions and reactions)
func (hf *HuggingFaceServiceClient) TriggerWebhook(payload map[string]any, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload send for this service")
	}
	if _, err := hf.cc.TriggerWebHook(context.Background(), &gRPCService.WebHookTriggerReq{ActionId: uint32(actionID), Payload: payloadBytes}); err != nil {
		return nil, err
	}
	return &IServ.WebHookResponseStatus{Description: "Webhook triggered"}, nil
}
