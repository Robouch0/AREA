//
// EPITECH PROJECT, 2024
// AREA
// File description:
// spotifyService
//

package spotify_server

import (
	"area/db"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"bytes"
	"cmp"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

type SpotifyService struct {
	tokenDb      *db.TokenDb
	spotifyDb    *db.SpotifyDB
	c            *cron.Cron
	reactService gRPCService.ReactionServiceClient

	gRPCService.UnimplementedSpotifyServiceServer
}

func NewSpotifyService() (*SpotifyService, error) {
	scheduler := cron.New()
	scheduler.Start()
	tokenDb, errT := db.InitTokenDb()
	spotifyDb, errS := db.InitSpotifyDb()
	if err := cmp.Or(errT, errS); err != nil {
		return nil, err
	}

	Spotify := &SpotifyService{
		tokenDb:      tokenDb,
		spotifyDb:    spotifyDb,
		c:            scheduler,
		reactService: nil,
	}

	Spotify.c.AddFunc("* * * * *", Spotify.checkVolumeNbr)
	Spotify.c.AddFunc("@every 0h03m00s", Spotify.checkIsPlaying)
	Spotify.c.AddFunc("@every 0h03m00s", Spotify.checkRepeatSong)
	Spotify.c.AddFunc("@every 0h03m00s", Spotify.checkShufflePlaylist)
	Spotify.c.AddFunc("@every 0h03m00s", Spotify.checkFollowersNbr)
	return Spotify, nil
}

func (spotify *SpotifyService) InitReactClient(conn *grpc.ClientConn) {
	spotify.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (spotify *SpotifyService) createNewSpotifyInfo(
	userID uint,
	actionID int,
	actionType models.SpotifyActionType,
	artist_id string,
	followers uint,
	volume uint,
) error {
	_, err := spotify.spotifyDb.InsertNewSpotify(&models.Spotify{
		ActionID:   uint(actionID),
		UserID:     userID,
		ActionType: actionType,
		Activated:  true,
		ArtistID:   artist_id,
		Followers:  followers,
		Volume:     volume,
	})
	return err
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

func (spot *SpotifyService) SetActivate(ctx context.Context, req *gRPCService.SetActivateSpotify) (*gRPCService.SetActivateSpotify, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "spotify")
	if err != nil {
		return nil, err
	}
	_, err = spot.spotifyDb.SetActivateByActionID(req.Activated, userID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}
