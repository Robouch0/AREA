//
// EPITECH PROJECT, 2024
// AREA
// File description:
// actions
//

package models

import "github.com/uptrace/bun"

type Reaction struct {
	Service      string         `json:"service"`
	Microservice string         `json:"microservice"`
	Ingredients  map[string]any `json:"ingredients"`
}

type Reactions struct {
	bun.BaseModel `bun:"table:reactions,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	AreaID        uint `json:"area_id"` // No anotation here !

	Reaction *Reaction `bun:"reaction" json:"reaction"`

	PrevOutput map[string]interface{} `bun:"prev_out,type:jsonb" json:"prev_out,type:jsonb"`
	// Repeat always field
}
