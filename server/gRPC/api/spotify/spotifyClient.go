//
// EPITECH PROJECT, 2024
// AREA
// File description:
// githubClient
//

package spotify

import (
	IServ "area/gRPC/api/serviceInterface"
	gRPCService "area/protogen/gRPC/proto"
	"context"
	"encoding/json"
	"errors"
	"log"

	"google.golang.org/grpc"
)

type SpotifyClient struct {
	MicroservicesLauncher *IServ.MicroserviceLauncher
	cc                    gRPCService.SpotifyServiceClient
}

func NewSpotifyClient(conn *grpc.ClientConn) *SpotifyClient {
	micros := &IServ.MicroserviceLauncher{}
	spotify := &SpotifyClient{MicroservicesLauncher: micros, cc: gRPCService.NewSpotifyServiceClient(conn)}
	(*spotify.MicroservicesLauncher)["stopSong"] = spotify.stopSong
	return spotify
}

func (spot *SpotifyClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Spotify",
		RefName: "spotify",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Stop the current song playing on the last device connected",
				RefName: "stopSong",
				Type:    "reaction",

				Ingredients: map[string]string{
				},
			},
		},
	}
	return status, nil
}


func (spot *SpotifyClient) SendAction(body map[string]any, actionId int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in spotify  service (Next will be things)")
}

func (spot *SpotifyClient) stopSong(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	_, err := json.Marshal(ingredients)
	if err != nil {
        log.Println("Ingredients problems", err)
		return nil, err
	}

	_, err = spot.cc.StopSong(context.Background(), &gRPCService.SpotifyStopInfo{})
	if err != nil {
        log.Println("Error when running stopSong service", err)
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song stopped"}, nil
}

func (spot *SpotifyClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*spot.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput)
	}
    log.Println(microservice)
	return nil, errors.New("No such microservice")
}
