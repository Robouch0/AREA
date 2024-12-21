//
// EPITECH PROJECT, 2024
// AREA
// File description:
// spotifyClient
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
	(*spotify.MicroservicesLauncher)["createPlaylist"] = spotify.createPlaylist
	(*spotify.MicroservicesLauncher)["nextSong"] = spotify.nextSong
	(*spotify.MicroservicesLauncher)["previousSong"] = spotify.previousSong
	(*spotify.MicroservicesLauncher)["setPlaybackVolume"] = spotify.setPlaybackVolume
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
            IServ.MicroserviceStatus{
                Name:    "Create a spotify playlist",
                RefName: "createPlaylist",
                Type:    "reaction",

                Ingredients: map[string]string{
                    "playlistName":    "string",
                    "playlistDescription":    "string",
                    "public": "string",
                },
            },
            IServ.MicroserviceStatus{
                Name:    "Launch the next song",
                RefName: "nextSong",
                Type:    "reaction",

                Ingredients: map[string]string{
                },
            },
            IServ.MicroserviceStatus{
                Name:    "Launch the previous song",
                RefName: "previousSong",
                Type:    "reaction",

                Ingredients: map[string]string{
                },
            },
            IServ.MicroserviceStatus{
                Name:    "Change the playback Volume",
                RefName: "setPlaybackVolume",
                Type:    "reaction",

                Ingredients: map[string]string{
                    "volume": "string",
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

func (spot *SpotifyClient) createPlaylist(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var createReq gRPCService.SpotifyCreatePlaylist
	err = json.Unmarshal(jsonString, &createReq)
	if err != nil {
		return nil, err
	}

	_, err = spot.cc.CreatePlaylist(context.Background(), &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Playlist created"}, nil
}

func (spot *SpotifyClient) nextSong(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	_, err := json.Marshal(ingredients)
	if err != nil {
        log.Println("Ingredients problems", err)
		return nil, err
	}

	_, err = spot.cc.NextSong(context.Background(), &gRPCService.SpotifyNextInfo{})
	if err != nil {
        log.Println("Error when running songNext service", err)
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song skipped"}, nil
}

func (spot *SpotifyClient) previousSong(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	_, err := json.Marshal(ingredients)
	if err != nil {
        log.Println("Ingredients problems", err)
		return nil, err
	}

	_, err = spot.cc.PreviousSong(context.Background(), &gRPCService.SpotifyPreviousInfo{})
	if err != nil {
        log.Println("Error when running songPrevious service", err)
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Go back to previous song"}, nil
}

func (spot *SpotifyClient) setPlaybackVolume(ingredients map[string]any, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
    jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var setPlaybackVolume gRPCService.SpotifySetPlaybackVolumeInfo
	err = json.Unmarshal(jsonString, &setPlaybackVolume)
	if err != nil {
		return nil, err
	}

	_, err = spot.cc.SetPlaybackVolume(context.Background(), &setPlaybackVolume)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Change the playback volume"}, nil
}


func (spot *SpotifyClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*spot.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput)
	}
    log.Println(microservice)
	return nil, errors.New("No such microservice")
}
