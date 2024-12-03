//
// EPITECH PROJECT, 2024
// AREA
// File description:
// OAuth
//

package models

import (
	"time"
)

type OAuthToken struct {
	ID           int64     `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	UserID       int64     `bun:"user_id" json:"user_id"`
	Provider     string    `bun:"provider" json:"provider"`
	AccessToken  string    `bun:"access_token" json:"access_token"`
	ExpiresAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
}