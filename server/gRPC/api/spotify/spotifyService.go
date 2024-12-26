//
// EPITECH PROJECT, 2024
// AREA
// File description:
// spotifyService
//

package spotify

import (
	"area/db"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type SpotifyService struct {
	tokenDb      *db.TokenDb
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedSpotifyServiceServer
}

func NewSpotifyService() (*SpotifyService, error) {
	tokenDb, err := db.InitTokenDb()

	return &SpotifyService{tokenDb: tokenDb, reactService: nil}, err
}

func (spot *SpotifyService) StopSong(ctx context.Context, req *gRPCService.SpotifyStopInfo) (*gRPCService.SpotifyStopInfo, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/pause")
	putRequest, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Println("Error when creating api call to spotify", err)
		return nil, err
	}
	putRequest.Header = http.Header{}
	putRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
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

func (spot *SpotifyService) CreatePlaylist(ctx context.Context, req *gRPCService.SpotifyCreatePlaylist) (*gRPCService.SpotifyCreatePlaylist, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}

	if req.PlaylistName == "" || req.PlaylistDescription == "" || req.Public == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	id, err := spot.GetUserID("Bearer " + tokenInfo.AccessToken)
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
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
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

func (spot *SpotifyService) NextSong(ctx context.Context, req *gRPCService.SpotifyNextInfo) (*gRPCService.SpotifyNextInfo, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/next")
	postRequest, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println("Error when creating api call to spotify", err)
		return nil, err
	}
	postRequest.Header = http.Header{}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
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

func (spot *SpotifyService) PreviousSong(ctx context.Context, req *gRPCService.SpotifyPreviousInfo) (*gRPCService.SpotifyPreviousInfo, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/previous")
	postRequest, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println("Error when creating api call to spotify", err)
		return nil, err
	}
	postRequest.Header = http.Header{}
	postRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
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

func (spot *SpotifyService) SetPlaybackVolume(ctx context.Context, req *gRPCService.SpotifySetPlaybackVolumeInfo) (*gRPCService.SpotifySetPlaybackVolumeInfo, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}

	if req.Volume == "" {
		log.Println("la", req.Volume)
		return nil, errors.New("Some required parameters are empty")
	}
	volume, err := strconv.ParseInt(req.Volume, 10, 32)
	if err != nil {
		log.Println("Error while parsing the int : ", err)
		return nil, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/volume?volume_percent=%d", volume)
	putRequest, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Println("Error when creating api call to spotify", err)
		return nil, err
	}
	putRequest.Header = http.Header{}
	putRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
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

func (spot *SpotifyService) LaunchSong(ctx context.Context, req *gRPCService.SpotifyLauchSongInfo) (*gRPCService.SpotifyLauchSongInfo, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}

	if req.SongUrl == "" || req.MillisecondsPosition == "" {
		return nil, errors.New("Some required parameters are empty")
	}
	milliseconds, err := strconv.ParseInt(req.MillisecondsPosition, 10, 32)
	if err != nil {
		log.Println("Error while parsing the int : ", err)
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
	putRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
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

func (spot *SpotifyService) AddSongToPlaylist(ctx context.Context, req *gRPCService.SpotifyAddSongToPlaylist) (*gRPCService.SpotifyAddSongToPlaylist, error) {
	tokenInfo, err := grpcutils.GetTokenByContext(ctx, spot.tokenDb, "SpotifyService", "spotify")
	if err != nil {
		return nil, err
	}
	//     if req.SongUrl == "" || req.MillisecondsPosition == ""{
	//         return nil, errors.New("Some required parameters are empty")
	//     }
	//     milliseconds, err := strconv.ParseInt(req.MillisecondsPosition, 10, 32)
	//     if err != nil {
	//         log.Println("Error while parsing the int : ", err)
	//         return nil, err
	//     }

	playlists, _ := spot.GetUserPlaylists("Bearer " + tokenInfo.AccessToken)

	foundID, _ := spot.FindPlayList(playlists, "My Playlist #2")
	log.Println(foundID)

	//     url := fmt.Sprintf("https://api.spotify.com/v1/me/player/play")
	//     putRequestBody := fmt.Sprintf(`{"uris": ["spotify:track:%s"], "position_ms":%d}`, req.SongUrl, milliseconds)
	//     putRequest, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(putRequestBody)))
	//     if err != nil {
	//         log.Println("Error when creating api call to spotify", err)
	//         return nil, err
	//     }
	//
	// 	putRequest.Header = http.Header{}
	//     putRequest.Header.Set("Authorization", "Bearer "+tokenInfo.AccessToken)
	//     putRequest.Header.Set("Content-Type", "application/json")
	//
	//     cli := &http.Client{}
	//     resp, err := cli.Do(putRequest)
	//     if err != nil {
	//         log.Println("Error when sending api call to spotify", err)
	//         return nil, err
	//     }
	//
	// 	if resp.StatusCode != 204 {
	// 	    log.Println("here", resp.Status)
	// 		return nil, errors.New(resp.Status)
	// 	}
	return req, nil
}
