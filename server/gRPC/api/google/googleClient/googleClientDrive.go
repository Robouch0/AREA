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
	conv_utils "area/utils/convUtils"
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

	return &IServ.ReactionResponseStatus{Description: res.Name, Datas: conv_utils.ConvertToMap[gRPCService.DeleteSharedDriveReq](res)}, nil
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

	return &IServ.ReactionResponseStatus{Description: res.NewName, Datas: conv_utils.ConvertToMap[gRPCService.UpdateSharedDriveReq](res)}, nil
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

	return &IServ.ReactionResponseStatus{Description: res.Name, Datas: conv_utils.ConvertToMap[gRPCService.CreateSharedDriveReq](res)}, nil
}
