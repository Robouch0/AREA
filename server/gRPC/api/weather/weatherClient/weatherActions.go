//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherActions
//

package weather_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (weather *WeatherClient) SendTemperatureAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.TempTriggerReq{
		Activated: true,
		ActionId:  int32(actionID),
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = weather.cc.NewTemperatureTrigger(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}

func (weather *WeatherClient) SendIsDayAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.IsDayTriggerReq{
		Activated: true,
		ActionId:  int32(actionID),
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = weather.cc.NewIsDayTrigger(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}
