//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherClient
//

package weather_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"errors"

	"google.golang.org/grpc"
)

type WeatherClient struct {
	ActionLauncher *IServ.ActionLauncher

	cc gRPCService.WeatherServiceClient
}

func NewWeatherClient(conn *grpc.ClientConn) *WeatherClient {
	actions := &IServ.ActionLauncher{}
	weather := &WeatherClient{ActionLauncher: actions, cc: gRPCService.NewWeatherServiceClient(conn)}

	(*weather.ActionLauncher)["temperatureExceed"] = weather.SendTemperatureAction
	(*weather.ActionLauncher)["dayChanged"] = weather.SendIsDayAction
	(*weather.ActionLauncher)["rainWeather"] = weather.SendRainAction
	(*weather.ActionLauncher)["snowWeather"] = weather.SendSnowAction

	// send a request to create a cron job and every 30min trigger it
	return weather
}

func (weather *WeatherClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*weather.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such microservice")
}

func (weather *WeatherClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,

	userID int,
) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No microservice Reaction yet")
}

func (weather *WeatherClient) TriggerWebhook(webhook *IServ.WebhookInfos, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}

func (weather *WeatherClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := weather.cc.SetActivate(ctx, &gRPCService.SetActivateWeather{
		ActionId:  uint32(id),
		Activated: activated,
	})
	if err != nil {
		return nil, err
	}
	return &IServ.SetActivatedResponseStatus{
		ActionID:    id,
		Description: "DateTime Deactivated",
	}, nil
}

func (weather *WeatherClient) DeleteArea(ID uint, userID uint) (*IServ.DeleteResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(int(userID))
	_, err := weather.cc.DeleteAction(ctx, &gRPCService.DeleteWeatherActionReq{
		ActionId: uint32(ID),
	})
	if err != nil {
		return nil, err
	}
	return &IServ.DeleteResponseStatus{ID: ID}, nil
}
