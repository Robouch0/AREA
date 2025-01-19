//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// spotify
//

package models

import "github.com/uptrace/bun"

type SpotifyActionType int

const (
	CheckVolume SpotifyActionType = iota
	CheckFollowers
	CheckShuffle
	CheckRepeat
	CheckPlaying
)

type Spotify struct {
	bun.BaseModel `bun:"table:spotify,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID uint `bun:"action_id" json:"action_id"`
	UserID   uint `bun:"user_id" json:"user_id"`

	ActionType SpotifyActionType `bun:"action_type" json:"action_type"`
	Activated  bool              `bun:"activated" json:"activated,omitempty"`

	ArtistID  string `bun:"artist_id" json:"artist_id,omitempty"`
	Followers uint   `bun:"followers" json:"followers,omitempty"`
	Volume    uint   `bun:"volume" json:"volume,omitempty"`
}
