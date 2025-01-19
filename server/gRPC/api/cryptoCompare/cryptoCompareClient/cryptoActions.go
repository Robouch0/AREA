//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoActions
//

package cryptocompareclient

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (crypto *CryptoClient) SendIsHigherAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.IsHigherThanTriggerReq{
		Activated: true,
		ActionId:  int32(actionID),
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = crypto.cc.IsHigherThanTrigger(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}

func (crypto *CryptoClient) SendIsLowerAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.IsLowerThanTriggerReq{
		Activated: true,
		ActionId:  int32(actionID),
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = crypto.cc.IsLowerThanTrigger(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}
