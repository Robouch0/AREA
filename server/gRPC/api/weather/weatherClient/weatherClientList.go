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

		Microservices: []IServ.MicroserviceStatus{
			{
				Name:    "Each hour, when temperature exceed a certain amount trigger happens",
				RefName: "temperatureExceed",
				Type:    "action",
				Ingredients: map[string]string{
					"temperature": "float",
					"timezone":    "string",
					"region":      "string",
				},
			},
			{
				Name:    "Each Hour, when the day condition (night or day changes)",
				RefName: "dayChanged",
				Type:    "action",
				Ingredients: map[string]string{
					"timezone": "string",
					"region":   "string",
				},
			},
		},
	}
	return status, nil
}
