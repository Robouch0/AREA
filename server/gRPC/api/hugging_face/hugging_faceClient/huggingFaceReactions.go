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
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (hf *HuggingFaceServiceClient) SendTextGenerationReaction(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := hf.cc.LaunchTextGeneration(ctx, &gRPCService.TextGenerationReq{Inputs: ingredients["inputs"].(string)})
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.GeneratedText}, nil
}

func (hf *HuggingFaceServiceClient) CreateRepositoryReaction(ingredients map[string]any, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: res.Name}, nil
}
