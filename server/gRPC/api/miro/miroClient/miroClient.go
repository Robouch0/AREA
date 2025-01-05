//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroClient
//

package miro_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MiroClient struct {
	ActionLauncher *IServ.ActionLauncher

	cc gRPCService.MiroServiceClient
}

func NewMiroClient(conn *grpc.ClientConn) *MiroClient {
	actions := &IServ.ActionLauncher{}
	miro := &MiroClient{ActionLauncher: actions, cc: gRPCService.NewMiroServiceClient(conn)}

	(*miro.ActionLauncher)["watchItemCreated"] = miro.createWebhookItemCreated
	return miro
}

func (miro *MiroClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*miro.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such microservice")
}

func (miro *MiroClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	return nil, status.Errorf(codes.Internal, "No set Activate yet")
}

func (miro *MiroClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	return nil, status.Errorf(codes.Internal, "No reaction yet")
}

func (miro *MiroClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	b, err := json.Marshal(webhook.Payload)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid payload sent")
	}
	if microservice == "itemCreated" {
		_, err = miro.cc.TriggerItemCreated(context.Background(), &gRPCService.ItemCreatedTriggerReq{Payload: b, ActionId: uint32(actionID)})
		if err != nil {
			return nil, err
		}
		return &IServ.WebHookResponseStatus{}, nil
	}
	return nil, status.Errorf(codes.NotFound, "Microservice: %v not found", microservice)
}
