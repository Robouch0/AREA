//
// EPITECH PROJECT, 2024
// AREA
// File description:
// spotifyClient
//

package spotify_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SpotifyClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	cc                    gRPCService.SpotifyServiceClient
}

func NewSpotifyClient(conn *grpc.ClientConn) *SpotifyClient {
	micros := &IServ.ReactionLauncher{}
	spotify := &SpotifyClient{MicroservicesLauncher: micros, cc: gRPCService.NewSpotifyServiceClient(conn)}
	(*spotify.MicroservicesLauncher)["stopSong"] = spotify.stopSong
	(*spotify.MicroservicesLauncher)["createPlaylist"] = spotify.createPlaylist
	(*spotify.MicroservicesLauncher)["nextSong"] = spotify.nextSong
	(*spotify.MicroservicesLauncher)["previousSong"] = spotify.previousSong
	(*spotify.MicroservicesLauncher)["setPlaybackVolume"] = spotify.setPlaybackVolume
	(*spotify.MicroservicesLauncher)["launchSong"] = spotify.launchSong
	(*spotify.MicroservicesLauncher)["addSongToPlaylist"] = spotify.addSongToPlaylist
	return spotify
}

func (spot *SpotifyClient) SendAction(_ models.AreaScenario, _, _ int) (*IServ.ActionResponseStatus, error) {
	return nil, errors.New("No action supported in spotify  service (Next will be things)")
}

func (spot *SpotifyClient) stopSong(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	_, err := json.Marshal(ingredients)
	if err != nil {
		log.Println("Ingredients problems", err)
		return nil, err
	}

	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spot.cc.StopSong(ctx, &gRPCService.SpotifyStopInfo{})
	if err != nil {
		log.Println("Error when running stopSong service", err)
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song stopped"}, nil
}

func (spot *SpotifyClient) createPlaylist(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var createReq gRPCService.SpotifyCreatePlaylist
	err = json.Unmarshal(jsonString, &createReq)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spot.cc.CreatePlaylist(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Playlist created"}, nil
}

func (spot *SpotifyClient) nextSong(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	_, err := json.Marshal(ingredients)
	if err != nil {
		log.Println("Ingredients problems", err)
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)

	_, err = spot.cc.NextSong(ctx, &gRPCService.SpotifyNextInfo{})
	if err != nil {
		log.Println("Error when running songNext service", err)
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song skipped"}, nil
}

func (spot *SpotifyClient) previousSong(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	_, err := json.Marshal(ingredients)
	if err != nil {
		log.Println("Ingredients problems", err)
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)

	_, err = spot.cc.PreviousSong(ctx, &gRPCService.SpotifyPreviousInfo{})
	if err != nil {
		log.Println("Error when running songPrevious service", err)
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Go back to previous song"}, nil
}

func (spot *SpotifyClient) setPlaybackVolume(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var setPlaybackVolume gRPCService.SpotifySetPlaybackVolumeInfo
	err = json.Unmarshal(jsonString, &setPlaybackVolume)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)

	_, err = spot.cc.SetPlaybackVolume(ctx, &setPlaybackVolume)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Change the playback volume"}, nil
}

func (spot *SpotifyClient) launchSong(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	jsonString, err := json.Marshal(ingredients)
	if err != nil {
		return nil, err
	}
	var launchSong gRPCService.SpotifyLauchSongInfo
	err = json.Unmarshal(jsonString, &launchSong)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)

	_, err = spot.cc.LaunchSong(ctx, &launchSong)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song launched"}, nil
}

func (spot *SpotifyClient) addSongToPlaylist(ingredients map[string]any, _ []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	// 	jsonString, err := json.Marshal(ingredients)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	var launchSong gRPCService.SpotifyAddSongToPlaylist
	// 	err = json.Unmarshal(jsonString, &launchSong)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	var addSong gRPCService.SpotifyAddSongToPlaylist
	_, err := spot.cc.AddSongToPlaylist(ctx, &addSong)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song added"}, nil
}

func (spot *SpotifyClient) TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*spot.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, prevOutput, userID)
	}
	log.Println(microservice)
	return nil, errors.New("No such microservice")
}

func (_ *SpotifyClient) TriggerWebhook(webhook *IServ.WebhookInfos, _ string, _ int) (*IServ.WebHookResponseStatus, error) {
	return &IServ.WebHookResponseStatus{}, nil
}

func (spot *SpotifyClient) DeactivateArea(microservice string, id uint, userID int) (*IServ.DeactivateResponseStatus, error) {
	return nil, status.Errorf(codes.Unavailable, "No action available yet for spotify")
}
