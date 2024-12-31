//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleDrive
//

package models

import "github.com/uptrace/bun"

type Drive struct {
	bun.BaseModel `bun:"table:drive,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	ChannelID string `bun:"channel_id" json:"channel_id,omitempty"`
}
