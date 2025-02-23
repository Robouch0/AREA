//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceClient
//

package huggingFace_client

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
	MicroservicesLauncher *IServ.ReactionLauncher
	ActionsLauncher       *IServ.ActionLauncher

	cc gRPCService.HuggingFaceServiceClient
}

// Alias for the webhook send function for HuggingFace
type WebHookRepoSendFunction = func(ctx context.Context, in *gRPCService.HFWebHookInfo, opts ...grpc.CallOption) (*gRPCService.HFWebHookInfo, error)

func NewHuggingFaceClient(conn *grpc.ClientConn) *HuggingFaceServiceClient {
	hf := &HuggingFaceServiceClient{
		MicroservicesLauncher: &IServ.ReactionLauncher{},
		ActionsLauncher:       &IServ.ActionLauncher{},
		cc:                    gRPCService.NewHuggingFaceServiceClient(conn),
	}
	(*hf.MicroservicesLauncher)["textGen"] = hf.SendTextGenerationReaction
	(*hf.MicroservicesLauncher)["createRepo"] = hf.CreateRepositoryReaction
	(*hf.MicroservicesLauncher)["deleteRepo"] = hf.DeleteRepositoryReaction
	(*hf.MicroservicesLauncher)["moveRepo"] = hf.MoveRepoReaction
	(*hf.MicroservicesLauncher)["changeVisibility"] = hf.ChangeRepoVisibilityReaction

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

func (hf *HuggingFaceServiceClient) TriggerReaction(ingredients map[string]any, microservice string, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*hf.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, userID)
	}
	return nil, errors.New("No such microservice")
}

// No need to check microservice here but later there will be a map for multiple type of actions (like actions and reactions)
func (hf *HuggingFaceServiceClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	payloadBytes, err := json.Marshal(webhook.Payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload send for this service")
	}
	if _, err := hf.cc.TriggerWebHook(context.Background(), &gRPCService.WebHookTriggerReq{ActionId: uint32(actionID), Payload: payloadBytes}); err != nil {
		return nil, err
	}
	return &IServ.WebHookResponseStatus{Description: "Webhook triggered"}, nil
}

func (hf *HuggingFaceServiceClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	return nil, nil
}

func (hf *HuggingFaceServiceClient) DeleteArea(ID uint, userID uint) (*IServ.DeleteResponseStatus, error) {
	return nil, status.Errorf(codes.Unavailable, "No Action for Discord Service yet")
}
