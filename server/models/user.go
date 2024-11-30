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
	Email         string `bun:"email" json:"email"`

	// Useful for log and security purposes
	CreatedAt time.Time `bun:"-" json:"-"`
	UpdatedAt time.Time `bun:"-" json:"-"`
}
