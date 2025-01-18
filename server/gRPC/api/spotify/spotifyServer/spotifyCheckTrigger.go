//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// spotifyCheckTrigger
//

package spotify_server

import (
	"area/models"
	service "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"log"
)

func (spotify *SpotifyService) checkVolumeNbr() {
	actions, err := spotify.spotifyDb.GetActionsByType(models.CheckVolume)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}

	for _, act := range *actions {
		tokenInfo, err := spotify.tokenDb.GetUserTokenByProvider(int64(act.UserID), "spotify")
		if err != nil {
			continue
		}
		resp, err := GetSpotifyInfos(tokenInfo.AccessToken)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.DeviceInfo.Volume > int32(act.Volume) {
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			spotify.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

func (spotify *SpotifyService) checkFollowersNbr() {
	actions, err := spotify.spotifyDb.GetActionsByType(models.CheckFollowers)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}
	for _, act := range *actions {
		tokenInfo, err := spotify.tokenDb.GetUserTokenByProvider(int64(act.UserID), "spotify")
		if err != nil {
			continue
		}
		resp, err := GetArtistInfos(act.ArtistID, tokenInfo.AccessToken)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.FollowerInfo.Followers > int32(act.Followers) {
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			spotify.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

func (spotify *SpotifyService) checkRepeatSong() {
	actions, err := spotify.spotifyDb.GetActionsByType(models.CheckRepeat)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}
	for _, act := range *actions {
		tokenInfo, err := spotify.tokenDb.GetUserTokenByProvider(int64(act.UserID), "spotify")
		if err != nil {
			continue
		}
		resp, err := GetSpotifyInfos(tokenInfo.AccessToken)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.RepeatState == "track" || resp.RepeatState == "context" {
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			spotify.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

func (spotify *SpotifyService) checkShufflePlaylist() {
	actions, err := spotify.spotifyDb.GetActionsByType(models.CheckShuffle)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}
	for _, act := range *actions {
		tokenInfo, err := spotify.tokenDb.GetUserTokenByProvider(int64(act.UserID), "spotify")
		if err != nil {
			continue
		}
		resp, err := GetSpotifyInfos(tokenInfo.AccessToken)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.ShuffleState == true {
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			spotify.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

func (spotify *SpotifyService) checkIsPlaying() {
	actions, err := spotify.spotifyDb.GetActionsByType(models.CheckPlaying)
	if err != nil {
		log.Println("Error while getting actions")
		return
	}
	for _, act := range *actions {
		tokenInfo, err := spotify.tokenDb.GetUserTokenByProvider(int64(act.UserID), "spotify")
		if err != nil {
			continue
		}
		resp, err := GetSpotifyInfos(tokenInfo.AccessToken)
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.SongPlaying == true {
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp)
			if err != nil {
				log.Println("Could not marshal crypto current response")
				continue
			}
			spotify.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}
