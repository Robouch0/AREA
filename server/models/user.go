//
// EPITECH PROJECT, 2024
// AREA
// File description:
// user
//

package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:user,alias:cs"`
	ID            uint   `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	FirstName     string `bun:"first_name" json:"first_name"`
	LastName      string `bun:"last_name" json:"last_name"`

	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`

	// Useful for log and security purposes
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
}

type OAuthToken struct {
	ID           int64     `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	UserID       int64     `bun:"user_id" json:"user_id"`
	Provider     string    `bun:"provider" json:"provider"`
	AccessToken  string    `bun:"access_token" json:"access_token"`
	ExpiresAt    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"-"`
}
