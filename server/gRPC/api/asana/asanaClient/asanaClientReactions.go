//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloClientReaction
//

package asana_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (asana *AsanaClient) createProject(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}

	var createReq gRPCService.CreateProjectReq
	err = json.Unmarshal(jsonString, &createReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := asana.cc.CreateProject(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Project created", Datas: conv_utils.ConvertToMap[gRPCService.CreateProjectResp](res)}, nil
}

func (asana *AsanaClient) createSection(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}

	var createReq gRPCService.CreateSectionReq
	err = json.Unmarshal(jsonString, &createReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := asana.cc.CreateSection(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Section created", Datas: conv_utils.ConvertToMap[gRPCService.CreateSectionResp](res)}, nil
}

func (asana *AsanaClient) createTask(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}

	var createReq gRPCService.CreateTaskReq
	err = json.Unmarshal(jsonString, &createReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := asana.cc.CreateTask(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Task created", Datas: conv_utils.ConvertToMap[gRPCService.CreateTaskResp](res)}, nil
}
