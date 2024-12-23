//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTime
//

package models

import "github.com/uptrace/bun"

type DateTime struct {
	bun.BaseModel `bun:"table:dateTime,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID uint `bun:"action_id" json:"action_id"`
	UserID   uint `bun:"user_id" json:"user_id"`

	Activated bool  `bun:"activated" json:"activated,omitempty"`
	Minutes   int32 `bun:"minutes" json:"minutes,omitempty"`
	Hours     int32 `bun:"hours" json:"hours,omitempty"`
	DayMonth  int64 `bun:"day_month" json:"day_month,omitempty"`
	Month     int64 `bun:"month" json:"month,omitempty"`
	DayWeek   int64 `bun:"day_week" json:"day_week,omitempty"`
}
