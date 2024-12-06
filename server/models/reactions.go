//
// EPITECH PROJECT, 2024
// AREA
// File description:
// actions
//

package models

import "github.com/uptrace/bun"

type Reactions struct {
	bun.BaseModel `bun:"table:user,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	AreaID        uint `bun:"area_id" json:"area_id"`

	Service      string `bun:"service" json:"service"`
	Microservice string `bun:"microservice" json:"microservice"`

	Ingredients map[string]interface{} `bun:"ingredients,type:jsonb" json:"ingredients,type:jsonb"`
	PrevOutput  map[string]interface{} `bun:"prev_out,type:jsonb" json:"prev_out,type:jsonb"`
}
