//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroServiceAction
//

package miro_server

import (
	miro_webhook "area/gRPC/api/miro/miroWebhook"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func enableToString(status bool) string {
	if status {
		return "enabled"
	} else {
		return "disabled"
	}
}

func (miro *MiroService) WatchItemCreated(ctx context.Context, req *gRPCService.ItemCreatedReq) (*gRPCService.ItemCreatedReq, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, miro.tokenDb, "MiroService", "miro")
	if err != nil {
		return nil, err
	}
	webhookURL, err := utils.GetEnvParameter("WEBHOOK_URL")
	if err != nil {
		return nil, err
	}
	callbackURL := fmt.Sprintf(webhookURL, "miro", "itemCreated", req.ActionId)

	res, err := miro_webhook.CreateWebhook(tokenInfo.AccessToken, &miro_webhook.CreateWebhookBody{
		BoardId:     req.BoardId,
		CallbackURL: callbackURL,
		Status:      enableToString(req.Status),
	})
	if err != nil {
		return nil, err
	}
	_, err = miro.miroDb.StoreNewMiro(&models.Miro{
		ActionID:  uint(req.ActionId),
		UserID:    uint(tokenInfo.UserID),
		Activated: false,
		WebhookID: res.ID,
		Type:      models.WCreate,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Database error while storing webhook: %v", err)
	}
	return req, nil
}

func (miro *MiroService) TriggerItemCreated(ctx context.Context, req *gRPCService.ItemCreatedTriggerReq) (*gRPCService.ItemCreatedResp, error) {
	var payload miro_webhook.MiroWebhookPayload
	if json.Unmarshal(req.Payload, &payload) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Payload received")
	}

	miroData, err := miro.miroDb.GetMiroByActionID(uint(req.ActionId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "No such datas for action id: %v", req.ActionId)
	}
	if miroData.Activated && payload.Type == string(miroData.Type) {
		ctx := grpcutils.CreateContextFromUserID(int(miroData.UserID))
		_, err := miro.reactService.LaunchReaction(
			ctx,
			&gRPCService.LaunchRequest{ActionId: int64(miroData.ActionID), PrevOutput: []byte(req.Payload)},
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return &gRPCService.ItemCreatedResp{
			Type:    payload.Type,
			Content: payload.Item.Data.Content,
		}, nil
	}
	return &gRPCService.ItemCreatedResp{
		Type:    "",
		Content: "",
	}, nil
}
