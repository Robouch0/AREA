//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceReactions
//

package huggingFace_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (hf *HuggingFaceServiceClient) SendTextGenerationReaction(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.LaunchTextGeneration(ctx, &gRPCService.TextGenerationReq{Inputs: ingredients["inputs"].(string)})
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.GeneratedText, Datas: conv_utils.ConvertToMap[gRPCService.TextGenerationRes](res)}, nil
}

func (hf *HuggingFaceServiceClient) CreateRepositoryReaction(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var createRepoReq gRPCService.CreateHFRepoReq
	err = json.Unmarshal(jsonString, &createRepoReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.CreateRepository(ctx, &createRepoReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Name, Datas: conv_utils.ConvertToMap[gRPCService.CreateHFRepoReq](res)}, nil
}

func (hf *HuggingFaceServiceClient) ChangeRepoVisibilityReaction(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var changeVisibilityReq gRPCService.ChangeHFRepoReq
	err = json.Unmarshal(jsonString, &changeVisibilityReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.ChangeRepoVisibility(ctx, &changeVisibilityReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Type, Datas: conv_utils.ConvertToMap[gRPCService.ChangeHFRepoReq](res)}, nil
}

func (hf *HuggingFaceServiceClient) MoveRepoReaction(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var moveReq gRPCService.MoveHFRepoReq
	err = json.Unmarshal(jsonString, &moveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.MoveRepo(ctx, &moveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.ToRepo, Datas: conv_utils.ConvertToMap[gRPCService.MoveHFRepoReq](res)}, nil
}

func (hf *HuggingFaceServiceClient) DeleteRepositoryReaction(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var deleteRepoReq gRPCService.DeleteHFRepoReq
	err = json.Unmarshal(jsonString, &deleteRepoReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.DeleteRepository(ctx, &deleteRepoReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Name, Datas: conv_utils.ConvertToMap[gRPCService.DeleteHFRepoReq](res)}, nil
}
