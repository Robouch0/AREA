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
	bun.BaseModel `bun:"table:customers,alias:cs"`
	ID            uint   `bun:"id,pk,autoincrement"`
	FirstName     string `bun:"first_name"`
	LastName      string `bun:"last_name"`
	Email         string `bun:"email"`

	// Useful for log and security purposes
	CreatedAt time.Time `bun:"-"`
	UpdatedAt time.Time `bun:"-"`
}
