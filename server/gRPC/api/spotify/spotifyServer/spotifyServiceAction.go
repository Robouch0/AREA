//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// spotifyServiceAction
//

package spotify_server

import (
	"context"
	"errors"
	"log"

	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
)

func (spotify *SpotifyService) CheckSongSoundVolume(ctx context.Context, req *gRPCService.SpotifyCheckVolume) (*gRPCService.SpotifyCheckVolume, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "SpotifyService")
	if err != nil {
		return nil, err
	}

	err = spotify.createNewSpotifyInfo(userID, int(req.ActionId), models.CheckVolume, "", 0, uint(req.Volume))
	if err != nil {
		return nil, err
	}
	log.Println("Spotify volume is being looked at")
	return req, nil
}

func (spotify *SpotifyService) CheckArtistFollowers(ctx context.Context, req *gRPCService.SpotifyCheckFollowers) (*gRPCService.SpotifyCheckFollowers, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "SpotifyService")
	if err != nil {
		return nil, err
	}

	if req.ArtistId == "" {
		return nil, errors.New("Some required parameters are empty")
	}

	err = spotify.createNewSpotifyInfo(userID, int(req.ActionId), models.CheckFollowers, req.ArtistId, uint(req.Followers), 0)
	if err != nil {
		return nil, err
	}
	log.Println("Spotify artist followers is being looked at")
	return req, nil
}

func (spotify *SpotifyService) CheckSongRepeat(ctx context.Context, req *gRPCService.SpotifyCheckRepeat) (*gRPCService.SpotifyCheckRepeat, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "SpotifyService")
	if err != nil {
		return nil, err
	}

	err = spotify.createNewSpotifyInfo(userID, int(req.ActionId), models.CheckRepeat, "", 0, 0)
	if err != nil {
		return nil, err
	}
	log.Println("Spotify song repeat is being looked at")
	return req, nil
}

func (spotify *SpotifyService) CheckPlaylistShuffle(ctx context.Context, req *gRPCService.SpotifyCheckShuffle) (*gRPCService.SpotifyCheckShuffle, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "SpotifyService")
	if err != nil {
		return nil, err
	}

	err = spotify.createNewSpotifyInfo(userID, int(req.ActionId), models.CheckShuffle, "", 0, 0)
	if err != nil {
		return nil, err
	}
	log.Println("Spotify playlist shuffle is being looked at")
	return req, nil
}

func (spotify *SpotifyService) CheckSongPlaylist(ctx context.Context, req *gRPCService.SpotifyCheckPlaying) (*gRPCService.SpotifyCheckPlaying, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "SpotifyService")
	if err != nil {
		return nil, err
	}

	err = spotify.createNewSpotifyInfo(userID, int(req.ActionId), models.CheckPlaying, "", 0, 0)
	if err != nil {
		return nil, err
	}
	log.Println("Spotify playing is being looked at")
	return req, nil
}
