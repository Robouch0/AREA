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

	Tokens []*OAuthToken `bun:"tokens,rel:has-many,join:id=user_id" json:"tokens"`
	Areas  []*Area       `bun:"areas,rel:has-many,join:id=user_id" json:"areas"`

	// Useful for log and security purposes
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
