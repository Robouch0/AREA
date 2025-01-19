//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// spotifyActions
//

package spotify_client

import (
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
)

func (spotify *SpotifyClient) SendCheckVolumeAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.SpotifyCheckVolume{
		ActionId: uint32(actionID),
		Activated: true,
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spotify.cc.CheckSongSoundVolume(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}

func (spotify *SpotifyClient) SendCheckFollowersAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.SpotifyCheckFollowers{
		ActionId: uint32(actionID),
		Activated: true,
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spotify.cc.CheckArtistFollowers(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}

func (spotify *SpotifyClient) SendCheckRepeatAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.SpotifyCheckRepeat{
		ActionId: uint32(actionID),
		Activated: true,
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spotify.cc.CheckSongRepeat(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}

func (spotify *SpotifyClient) SendCheckShuffleAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.SpotifyCheckShuffle{
		ActionId: uint32(actionID),
		Activated: true,
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spotify.cc.CheckPlaylistShuffle(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}

func (spotify *SpotifyClient) SendCheckPlayingAction(scenario models.AreaScenario, actionID, userID int) (*IServ.ActionResponseStatus, error) {
	wReqBytes, err := json.Marshal(scenario.Action.Ingredients)
	if err != nil {
		return nil, err
	}

	wRequest := gRPCService.SpotifyCheckPlaying{
		ActionId: uint32(actionID),
		Activated: true,
	}
	err = json.Unmarshal(wReqBytes, &wRequest)
	if err != nil {
		return nil, err
	}
	ctx := grpcutils.CreateContextFromUserID(userID)
	_, err = spotify.cc.CheckSongPlaylist(ctx, &wRequest)
	if err != nil {
		return nil, err
	}
	return &IServ.ActionResponseStatus{
		Description: "Done",
		ActionID:    actionID,
	}, nil
}
