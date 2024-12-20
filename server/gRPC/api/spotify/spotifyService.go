//
// EPITECH PROJECT, 2024
// AREA
// File description:
// spotifyService
//

package spotify

import (
	gRPCService "area/protogen/gRPC/proto"
	"area/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type SpotifyService struct {
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedSpotifyServiceServer
}

func NewSpotifyService() SpotifyService {
	return SpotifyService{reactService: nil}
}

func (spot *SpotifyService) StopSong(_ context.Context, req *gRPCService.SpotifyStopInfo) (*gRPCService.SpotifyStopInfo, error) {
	bearerTok, err := utils.GetEnvParameterToBearer("API_SPOTIFY")
    if err != nil {
        log.Println("No api bearer SPOTIFY : ", err)
        return nil, err
    }
    url := fmt.Sprintf("https://api.spotify.com/v1/me/player/pause")
    putRequest, err := http.NewRequest("PUT", url, nil)
    if err != nil {
        log.Println("Error when creating api call to spotify", err)
        return nil, err
    }
	putRequest.Header = utils.GetDefaultHTTPHeader(bearerTok)
    cli := &http.Client{}
    resp, err := cli.Do(putRequest)
    if err != nil {
        log.Println("Error when sending api call to spotify", err)
        return nil, err
    }

	if resp.Status != "200 OK" {
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body) // Do something with it
	return req, nil
}
