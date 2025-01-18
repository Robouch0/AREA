//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroClientReactions
//

package miro_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (miro *MiroClient) createStickyNote(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var miroReq gRPCService.CreateStickyNoteReq
	err = json.Unmarshal(jsonString, &miroReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := miro.cc.CreateStickyNote(ctx, &miroReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ReactionResponseStatus{Description: "Done", Datas: conv_utils.ConvertToMap[gRPCService.CreateStickyNoteResp](res)}, nil
}

func (miro *MiroClient) createTextNote(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var miroReq gRPCService.CreateTextReq
	err = json.Unmarshal(jsonString, &miroReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := miro.cc.CreateTextItem(ctx, &miroReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ReactionResponseStatus{Description: "Done", Datas: conv_utils.ConvertToMap[gRPCService.CreateTextResp](res)}, nil
}

func (miro *MiroClient) createCardNote(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var miroReq gRPCService.CreateCardReq
	err = json.Unmarshal(jsonString, &miroReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := miro.cc.CreateCardItem(ctx, &miroReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &IServ.ReactionResponseStatus{Description: "Done", Datas: conv_utils.ConvertToMap[gRPCService.CreateCardResp](res)}, nil
}
