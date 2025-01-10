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

	Timezone  string  `bun:"timezone" json:"timezone,omitempty"`
	Latitude  float64 `bun:"latitude" json:"latitude,omitempty"`
	Longitude float64 `bun:"longitude" json:"longitude,omitempty"`

	Rain     float64 `bun:"rain" json:"rain,omitempty"`
	SnowFall float64 `bun:"snowfall" json:"snowfall,omitempty"`

	IsDay int `bun:"is_day" json:"is_day,omitempty"`
}
