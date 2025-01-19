//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// spotifyTypes
//

package spotify_server

type SpotifyDeviceInfo struct {
	Volume int32 `json:"volume_percent,omitempty"`
}

type SpotifyFollowersInfo struct {
	Followers int32 `json:"total,omitempty"`
}

type SpotifyArtistAPIResponseBody struct {
	FollowerInfo SpotifyFollowersInfo `json:"followers"`
}

type SpotifyInfoAPIResponseBody struct {
	RepeatState  string            `json:"repeat_state,omitempty"`
	ShuffleState bool              `json:"shuffle_state,omitempty"`
	DeviceInfo   SpotifyDeviceInfo `json:"device,omitempty"`
	SongPlaying  bool              `json:"is_playing,omitempty"`
}
