//
// EPITECH PROJECT, 2024
// AREA
// File description:
// area
//

package models

import "github.com/uptrace/bun"

type Area struct {
	bun.BaseModel `bun:"table:areas,alias:area"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	UserID        uint `bun:"user_id" json:"user_id"`

	Action    *Actions     `bun:"action,rel:has-one,join:id=area_id" json:"action"`
	Reactions []*Reactions `bun:"reactions,rel:has-many,join:id=area_id" json:"reactions"`

	Activated bool `bun:"activated" json:"activated"`
	OneShot   bool `bun:"one_shot" json:"one_shot"`
}

type AreaScenario struct {
	Action    Action     `json:"action"`
	Reactions []Reaction `json:"reactions"`
}
