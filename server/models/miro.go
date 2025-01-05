//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miro
//

package models

import "github.com/uptrace/bun"

type WMiroType string

const (
	WCreate WMiroType = "create"
	WDelete           = "delete"
	WUpdate           = "update"
)

type Miro struct {
	bun.BaseModel `bun:"table:miro,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	WebhookID string    `bun:"webhook_id" json:"webhook_id,omitempty"` /* Useful for the activate/deactivate */
	Type      WMiroType `bun:"type" json:"type,omitempty"`
}
