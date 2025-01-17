//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleClientDrive
//

package google_client

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (google *GoogleClient) deleteDrive(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.DeleteSharedDriveReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.DeleteSharedDrive(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Name}, nil
}

func (google *GoogleClient) updateDrive(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.UpdateSharedDriveReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.UpdateSharedDrive(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.NewName}, nil
}

func (google *GoogleClient) createDrive(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var driveReq gRPCService.CreateSharedDriveReq
	err = json.Unmarshal(jsonString, &driveReq)
	if err != nil {
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	res, err := google.cc.CreateSharedDrive(ctx, &driveReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: res.Name}, nil
}
