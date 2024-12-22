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
	"bytes"
	"net/http"
	"strconv"
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
	log.Println(bearerTok)
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
	putRequest.Header = http.Header{}
    putRequest.Header.Set("Authorization", bearerTok)
    putRequest.Header.Set("Content-Type", "application/json")

    cli := &http.Client{}
    resp, err := cli.Do(putRequest)
    if err != nil {
        log.Println("Error when sending api call to spotify", err)
        return nil, err
    }

	if resp.StatusCode != 200 {
	    log.Println("here", resp.Status)
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body)
	return req, nil
}


func (spot *SpotifyService) CreatePlaylist(_ context.Context, req *gRPCService.SpotifyCreatePlaylist) (*gRPCService.SpotifyCreatePlaylist, error) {
	if req.PlaylistName == "" || req.PlaylistDescription == "" || req.Public == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	bearerTok, err := utils.GetEnvParameterToBearer("API_SPOTIFY")
	if err != nil {
		log.Println("No api bearer SPOTIFY: ", err)
		return nil, err
	}

	id, err := spot.GetUserID(bearerTok)
	if err != nil {
		log.Println("Error getting user ID: ", err)
		return nil, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", id)

	postRequestBody := fmt.Sprintf(`{"name": "%s", "description": "%s", "public": "%s"}`, req.PlaylistName, req.PlaylistDescription, req.Public)
	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(postRequestBody)))
	if err != nil {
		log.Println("Error when creating API call to Spotify", err)
		return nil, err
	}

	postRequest.Header = http.Header{}
	postRequest.Header.Set("Authorization", bearerTok)
	postRequest.Header.Set("Content-Type", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(postRequest)
	if err != nil {
		log.Println("Error when sending API call to Spotify", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Error creating playlist:", resp.Status)
		return nil, errors.New(resp.Status)
	}

	log.Println("Playlist created")
	return req, nil
}


func (spot *SpotifyService) NextSong(_ context.Context, req *gRPCService.SpotifyNextInfo) (*gRPCService.SpotifyNextInfo, error) {
	bearerTok, err := utils.GetEnvParameterToBearer("API_SPOTIFY")
	log.Println(bearerTok)
    if err != nil {
        log.Println("No api bearer SPOTIFY : ", err)
        return nil, err
    }
    url := fmt.Sprintf("https://api.spotify.com/v1/me/player/next")
    postRequest, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Println("Error when creating api call to spotify", err)
        return nil, err
    }
	postRequest.Header = http.Header{}
    postRequest.Header.Set("Authorization", bearerTok)
    postRequest.Header.Set("Content-Type", "application/json")

    cli := &http.Client{}
    resp, err := cli.Do(postRequest)
    if err != nil {
        log.Println("Error when sending api call to spotify", err)
        return nil, err
    }

	if resp.StatusCode != 200 {
	    log.Println("here", resp.Status)
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body)
	return req, nil
}


func (spot *SpotifyService) PreviousSong(_ context.Context, req *gRPCService.SpotifyPreviousInfo) (*gRPCService.SpotifyPreviousInfo, error) {
	bearerTok, err := utils.GetEnvParameterToBearer("API_SPOTIFY")
	log.Println(bearerTok)
    if err != nil {
        log.Println("No api bearer SPOTIFY : ", err)
        return nil, err
    }
    url := fmt.Sprintf("https://api.spotify.com/v1/me/player/previous")
    postRequest, err := http.NewRequest("POST", url, nil)
    if err != nil {
        log.Println("Error when creating api call to spotify", err)
        return nil, err
    }
	postRequest.Header = http.Header{}
    postRequest.Header.Set("Authorization", bearerTok)
    postRequest.Header.Set("Content-Type", "application/json")

    cli := &http.Client{}
    resp, err := cli.Do(postRequest)
    if err != nil {
        log.Println("Error when sending api call to spotify", err)
        return nil, err
    }

	if resp.StatusCode != 200 {
	    log.Println("here", resp.Status)
		return nil, errors.New(resp.Status)
	}
	log.Println("Here: ", resp.Body)
	return req, nil
}

func (spot *SpotifyService) SetPlaybackVolume(_ context.Context, req *gRPCService.SpotifySetPlaybackVolumeInfo) (*gRPCService.SpotifySetPlaybackVolumeInfo, error) {
    if req.Volume == "" {
        log.Println("la", req.Volume)
        return nil, errors.New("Some required parameters are empty")
    }
    volume, err := strconv.ParseInt(req.Volume, 10, 32)
    if err != nil {
        log.Println("Error while parsing the int : ", err)
        return nil, err
    }

	bearerTok, err := utils.GetEnvParameterToBearer("API_SPOTIFY")
	log.Println(bearerTok)
    if err != nil {
        log.Println("No api bearer SPOTIFY : ", err)
        return nil, err
    }
    url := fmt.Sprintf("https://api.spotify.com/v1/me/player/volume?volume_percent=%d", volume)
    putRequest, err := http.NewRequest("PUT", url, nil)
    if err != nil {
        log.Println("Error when creating api call to spotify", err)
        return nil, err
    }
	putRequest.Header = http.Header{}
    putRequest.Header.Set("Authorization", bearerTok)
    putRequest.Header.Set("Content-Type", "application/json")

    cli := &http.Client{}
    resp, err := cli.Do(putRequest)
    if err != nil {
        log.Println("Error when sending api call to spotify", err)
        return nil, err
    }

	if resp.StatusCode != 204 {
	    log.Println("here", resp.Status)
		return nil, errors.New(resp.Status)
	}
	return req, nil
}

func (spot *SpotifyService) LaunchSong(_ context.Context, req *gRPCService.SpotifyLauchSongInfo) (*gRPCService.SpotifyLauchSongInfo, error) {
    if req.SongUrl == "" || req.MillisecondsPosition == ""{
        return nil, errors.New("Some required parameters are empty")
    }
    milliseconds, err := strconv.ParseInt(req.MillisecondsPosition, 10, 32)
    if err != nil {
        log.Println("Error while parsing the int : ", err)
        return nil, err
    }

	bearerTok, err := utils.GetEnvParameterToBearer("API_SPOTIFY")
	log.Println(bearerTok)
    if err != nil {
        log.Println("No api bearer SPOTIFY : ", err)
        return nil, err
    }
    url := fmt.Sprintf("https://api.spotify.com/v1/me/player/play")
    putRequestBody := fmt.Sprintf(`{"uris": ["spotify:track:%s"], "position_ms":%d}`, req.SongUrl, milliseconds)
    putRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(putRequestBody)))
    if err != nil {
        log.Println("Error when creating api call to spotify", err)
        return nil, err
    }

	putRequest.Header = http.Header{}
    putRequest.Header.Set("Authorization", bearerTok)
    putRequest.Header.Set("Content-Type", "application/json")

    cli := &http.Client{}
    resp, err := cli.Do(putRequest)
    if err != nil {
        log.Println("Error when sending api call to spotify", err)
        return nil, err
    }

	if resp.StatusCode != 204 {
	    log.Println("here", resp.Status)
		return nil, errors.New(resp.Status)
	}
	return req, nil
}
