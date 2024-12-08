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
	bun.BaseModel `bun:"table:actions,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	AreaID        uint `json:"area_id"` // No anotation here !

	Action *Action `bun:"action" json:"action"`
}
