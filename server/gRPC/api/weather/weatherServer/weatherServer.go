//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherServer
//

package weather_server

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	"context"

	"google.golang.org/grpc"
)

type WeatherService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedWeatherServiceServer
}

func NewWeatherService() (*WeatherService, error) {
	tokenDb, err := db.InitTokenDb()
	if err != nil {
		return nil, err
	}

	return &WeatherService{tokenDb: tokenDb, reactService: nil}, err
}

func (weather *WeatherService) InitReactClient(conn *grpc.ClientConn) {
	weather.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (weather *WeatherService) NewTemperatureTrigger(ctx context.Context, req *gRPCService.TempTriggerReq) (*gRPCService.TempTriggerReq, error) {
	return nil, nil
}
