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

	// send a request to create a cron job and every 30min trigger it
	return weather
}

func (weather *WeatherClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No microservice Action yet")
}

func (weather *WeatherClient) TriggerReaction(
	ingredients map[string]any,
	microservice string,
	prevOutput []byte,
	userID int,
) (*IServ.ReactionResponseStatus, error) {
	return nil, errors.New("No microservice Reaction yet")
}

func (weather *WeatherClient) TriggerWebhook(payload map[string]any, microservice string, actionID int) (*IServ.WebHookResponseStatus, error) {
	return nil, errors.New("No microservice TriggerWebhook yet")
}
