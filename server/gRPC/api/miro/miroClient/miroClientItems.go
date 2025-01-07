//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroClientItems
//

package miro_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (miro *MiroClient) createWebhookItemCreated(scenario models.AreaScenario, actionID int, userID int) (*IServ.ActionResponseStatus, error) {
	jsonString, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}
	var miroReq gRPCService.ItemCreatedReq
	err = json.Unmarshal(jsonString, &miroReq)
	if err != nil {
		return nil, err
	}
	miroReq.ActionId = uint32(actionID)
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = miro.cc.WatchItemCreated(ctx, &miroReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ActionResponseStatus{Description: "Done", ActionID: actionID}, nil
}
