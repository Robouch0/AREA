//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTime
//

package models

import "github.com/uptrace/bun"

type DateTime struct {
	bun.BaseModel `bun:"table:user,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	ActionID      uint `bun:"action_id" json:"action_id"`
}
