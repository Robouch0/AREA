//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleClientsFiles
//

package google_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (google *GoogleClient) createEmptyFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.CreateEmptyFileReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.CreateEmptyFile(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Description}, nil
}

func (google *GoogleClient) deleteFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.DeleteFileReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.DeleteFile(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.FileName}, nil
}

func (google *GoogleClient) updateFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.UpdateFileMetaReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.UpdateFileMetadata(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.NewFileName}, nil
}

func (google *GoogleClient) copyFile(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.CopyFileReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.CopyFile(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.DestFileName}, nil
}
