//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherTypes
//

package weather_server

// Coordinate used as map for region supported
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Config

type WeatherConfig struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Current  string `json:"current,omitempty"`
	Timezone string `json:"timezone"`
}

// Body Response

type CurrentInformations struct {
	Time          string  `json:"time,omitempty"`
	Interval      float64 `json:"interval,omitempty"`
	Temperature2m float64 `json:"temperature_2m,omitempty"`
	IsDay         int     `json:"is_day,omitempty"`
}

type WeatherAPIResponseBody struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Timezone  string  `json:"timezone,omitempty"`

	// Add current units maybe
	Current CurrentInformations `json:"current"`
}
