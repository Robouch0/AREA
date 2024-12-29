//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherCheckTrigger
//

package weather_server

import (
	"area/models"
	service "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"encoding/json"
	"log"
)

func (weather *WeatherService) checkTemperature() {
	actions, err := weather.weatherDb.GetActionsByType(models.TemperatureExceed)
	if err != nil {
		log.Println("Error while loading temperatures")
		return
	}
	for _, act := range *actions {
		resp, err := GetCurrentWeather(&WeatherConfig{
			Latitude:  act.Latitude,
			Longitude: act.Longitude,
			Current:   "temperature_2m",
			Timezone:  act.Timezone,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.Current.Temperature2m > act.Temperature {
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp.Current)
			if err != nil {
				log.Println("Could not marshal weather current response")
				continue
			}
			weather.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}

func (weather *WeatherService) checkDayCondition() {
	actions, err := weather.weatherDb.GetActionsByType(models.TemperatureExceed)
	if err != nil {
		log.Println("Error while loading temperatures")
		return
	}
	for _, act := range *actions {
		resp, err := GetCurrentWeather(&WeatherConfig{
			Latitude:  act.Latitude,
			Longitude: act.Longitude,
			Current:   "is_day",
			Timezone:  act.Timezone,
		})
		if err != nil {
			log.Println(err)
			continue
		}
		if resp.Current.IsDay != act.IsDay {
			weather.weatherDb.UpdateUserIsDay(int(act.ActionID), act.IsDay)
			ctx := grpcutils.CreateContextFromUserID(int(act.UserID))
			b, err := json.Marshal(&resp.Current)
			if err != nil {
				log.Println("Could not marshal weather current response")
				continue
			}
			weather.reactService.LaunchReaction(ctx, &service.LaunchRequest{
				ActionId:   int64(act.ActionID),
				PrevOutput: b,
			})
		}
	}
}
