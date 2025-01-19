//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherClientList
//

package weather_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (weather *WeatherClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Weather",
		RefName: "weather",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Each hour, when temperature exceed a certain amount trigger happens",
				RefName: "temperatureExceed",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"temperature": {
						Value:       0.0,
						Type:        "float",
						Description: "Temperature reference that will be tracked",
						Required:    true,
					},
					"timezone": {
						Value:       "Europe/Paris",
						Type:        "string",
						Description: "Timezone that will be tracked",
						Required:    true,
					},
					"region": {
						Value:       "France",
						Type:        "string",
						Description: "Region that will be tracked",
						Required:    true,
					},
				},
				PipelineAvailable: []string{"time", "interval", "temperature_2m", "is_day", "rain", "snowfall"},
			},
			{
				Name:    "Each Hour, when the day condition (night or day changes)",
				RefName: "dayChanged",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"timezone": {
						Value:       "Europe/Paris",
						Type:        "string",
						Description: "Timezone that will be tracked",
						Required:    true,
					},
					"region": {
						Value:       "France",
						Type:        "string",
						Description: "Region that will be tracked",
						Required:    true,
					},
				},
				PipelineAvailable: []string{"time", "interval", "temperature_2m", "is_day", "rain", "snowfall"},
			},
			{
				Name:    "Each Hour, when the current weather condition is rain",
				RefName: "rainWeather",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"timezone": {
						Value:       "Europe/Paris",
						Type:        "string",
						Description: "Timezone that will be tracked",
						Required:    true,
					},
					"region": {
						Value:       "France",
						Type:        "string",
						Description: "Region that will be tracked",
						Required:    true,
					},
				},
				PipelineAvailable: []string{"time", "interval", "temperature_2m", "is_day", "rain", "snowfall"},
			},
			{
				Name:    "Each Hour, when the current weather condition is snow",
				RefName: "snowWeather",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"timezone": {
						Value:       "Europe/Paris",
						Type:        "string",
						Description: "Timezone that will be tracked",
						Required:    true,
					},
					"region": {
						Value:       "France",
						Type:        "string",
						Description: "Region that will be tracked",
						Required:    true,
					},
				},
				PipelineAvailable: []string{"time", "interval", "temperature_2m", "is_day", "rain", "snowfall"},
			},
		},
	}
	return status, nil
}
