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

		Microservices: []IServ.MicroserviceStatus{},
	}
	return status, nil
}
