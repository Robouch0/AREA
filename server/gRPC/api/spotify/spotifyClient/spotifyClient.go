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
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"errors"
	"log"

	"google.golang.org/grpc"
)

type SpotifyClient struct {
	MicroservicesLauncher *IServ.ReactionLauncher
	ActionLauncher        *IServ.ActionLauncher

	cc                    gRPCService.SpotifyServiceClient
}

func NewSpotifyClient(conn *grpc.ClientConn) *SpotifyClient {
	micros := &IServ.ReactionLauncher{}
	actions := &IServ.ActionLauncher{}
	spotify := &SpotifyClient{MicroservicesLauncher: micros, ActionLauncher: actions, cc: gRPCService.NewSpotifyServiceClient(conn)}
	(*spotify.MicroservicesLauncher)["stopSong"] = spotify.stopSong
	(*spotify.MicroservicesLauncher)["createPlaylist"] = spotify.createPlaylist
	(*spotify.MicroservicesLauncher)["nextSong"] = spotify.nextSong
	(*spotify.MicroservicesLauncher)["previousSong"] = spotify.previousSong
	(*spotify.MicroservicesLauncher)["setPlaybackVolume"] = spotify.setPlaybackVolume
	(*spotify.MicroservicesLauncher)["launchSong"] = spotify.launchSong
	(*spotify.MicroservicesLauncher)["addSongToPlaylist"] = spotify.addSongToPlaylist

	(*spotify.ActionLauncher)["checkFollowers"] = spotify.SendCheckFollowersAction
	(*spotify.ActionLauncher)["checkVolume"] = spotify.SendCheckVolumeAction
	(*spotify.ActionLauncher)["checkRepeat"] = spotify.SendCheckRepeatAction
	(*spotify.ActionLauncher)["checkShuffle"] = spotify.SendCheckShuffleAction
	(*spotify.ActionLauncher)["checkPlaying"] = spotify.SendCheckPlayingAction

	return spotify
}

func (spot *SpotifyClient) SendAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	if micro, ok := (*spot.ActionLauncher)[scenario.Action.Microservice]; ok {
		return micro(scenario, actionID, userID)
	}
	return nil, errors.New("No such microservice")
}

func (spot *SpotifyClient) stopSong(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: "Song stopped", Datas: map[string]any{}}, nil
}

func (spot *SpotifyClient) createPlaylist(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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
	res, err := spot.cc.CreatePlaylist(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Playlist created", Datas: conv_utils.ConvertToMap[gRPCService.SpotifyCreatePlaylist](res)}, nil
}

func (spot *SpotifyClient) nextSong(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: "Song skipped", Datas: map[string]any{}}, nil
}

func (spot *SpotifyClient) previousSong(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: "Go back to previous song", Datas: map[string]any{}}, nil
}

func (spot *SpotifyClient) setPlaybackVolume(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	res, err := spot.cc.SetPlaybackVolume(ctx, &setPlaybackVolume)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Change the playback volume", Datas: conv_utils.ConvertToMap[gRPCService.SpotifySetPlaybackVolumeInfo](res)}, nil
}

func (spot *SpotifyClient) launchSong(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	res, err := spot.cc.LaunchSong(ctx, &launchSong)
	if err != nil {
		return nil, err
	}

	return &IServ.ReactionResponseStatus{Description: "Song launched", Datas: conv_utils.ConvertToMap[gRPCService.SpotifyLauchSongInfo](res)}, nil
}

func (spot *SpotifyClient) addSongToPlaylist(ingredients map[string]any, userID int) (*IServ.ReactionResponseStatus, error) {
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

	return &IServ.ReactionResponseStatus{Description: "Song added", Datas: map[string]any{}}, nil // Not used
}

func (spot *SpotifyClient) TriggerReaction(ingredients map[string]any, microservice string, userID int) (*IServ.ReactionResponseStatus, error) {
	if micro, ok := (*spot.MicroservicesLauncher)[microservice]; ok {
		return micro(ingredients, userID)
	}
	log.Println(microservice)
	return nil, errors.New("No such microservice")
}

func (_ *SpotifyClient) TriggerWebhook(webhook *IServ.WebhookInfos, _ string, _ int) (*IServ.WebHookResponseStatus, error) {
	return &IServ.WebHookResponseStatus{}, nil
}

func (spot *SpotifyClient) SetActivate(microservice string, id uint, userID int, activated bool) (*IServ.SetActivatedResponseStatus, error) {
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err := spot.cc.SetActivate(ctx, &gRPCService.SetActivateSpotify{
		ActionId:  uint32(id),
		Activated: activated,
	})
	if err != nil {
		return nil, err
	}
	return &IServ.SetActivatedResponseStatus{
		ActionID:    id,
		Description: "Spotify Deactivated",
	}, nil
}
