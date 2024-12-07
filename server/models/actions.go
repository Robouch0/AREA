//
// EPITECH PROJECT, 2024
// AREA
// File description:
// actions
//

package models

import "github.com/uptrace/bun"

type Action struct {
	Service      string         `json:"service"`
	Microservice string         `json:"microservice"`
	Ingredients  map[string]any `json:"ingredients"`
}

type Actions struct {
	bun.BaseModel `bun:"table:user,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	AreaID        uint `bun:"area_id" json:"area_id"`

	Action Reaction `bun:"action" json:"action"`
}
