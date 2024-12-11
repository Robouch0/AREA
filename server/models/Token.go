//
// EPITECH PROJECT, 2024
// AREA
// File description:
// Token
//

package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Token struct {
	bun.BaseModel `bun:"table:tokens,alias:cs"`
	ID           int64     `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	UserID       int64     `bun:"user_id" json:"user_id"`
	User   *User `bun:"user,rel:belongs-to,join:user_id=id" json:"user"`

	Provider     string    `bun:"provider" json:"provider"`
	AccessToken  string    `bun:"access_token" json:"access_token"`
	ExpiresAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
}
