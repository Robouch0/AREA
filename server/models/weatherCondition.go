//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherCondition
//

package models

import "github.com/uptrace/bun"

type WeatherActionType int

const (
	TemperatureExceed WeatherActionType = iota
	DayCondition
)

type WeatherCondition struct {
	bun.BaseModel `bun:"table:weatherCondition,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID uint `bun:"action_id" json:"action_id"`
	UserID   uint `bun:"user_id" json:"user_id"`

	ActionType WeatherActionType `bun:"action_type" json:"action_type"`
	Activated  bool              `bun:"activated" json:"activated,omitempty"`

	Temperature        float64 `bun:"temperature" json:"temperature,omitempty"`
	TemperatureMetrics string  `bun:"temperature_metrics" json:"temperature_metrics,omitempty"` /* Celcius or Fahrenheit */

	IsDay bool `bun:"is_day" json:"is_day,omitempty"`
}
