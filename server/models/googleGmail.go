//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleGmail
//

package models

import "github.com/uptrace/bun"

type Gmail struct {
	bun.BaseModel `bun:"table:gmail,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	HistoryID   string `bun:"history_id" json:"historyId"`
	EmailAdress string `bun:"email_address" json:"emailAddress"`
	FirstTime   bool   `bun:"first_time" json:"first_time,omitempty"`
}
