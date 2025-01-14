//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoCompare
//

package models

import "github.com/uptrace/bun"

type CryptoActionType int

const (
	IsHigher CryptoActionType = iota
	IsLower
)

type CryptoCompare struct {
	bun.BaseModel `bun:"table:cryptoCompare,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID uint `bun:"action_id" json:"action_id"`
	UserID   uint `bun:"user_id" json:"user_id"`

	ActionType CryptoActionType `bun:"action_type" json:"action_type"`
	Activated  bool             `bun:"activated" json:"activated,omitempty"`

	CryptoCurrency string `bun:"crypto_currency" json:"crypto_currency"`
	Currency       string `bun:"currency" json:"currency"`
	Threshold      uint   `bun:"threshold" json:"threshold"`
}
