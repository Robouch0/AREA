//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherServer
//

package weather_server

import (
	"area/db"
	"area/models"
	gRPCService "area/protogen/gRPC/proto"
	grpcutils "area/utils/grpcUtils"
	"cmp"
	"context"
	"log"

	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WeatherService struct {
	tokenDb         *db.TokenDb
	weatherDb       *db.WeatherConditionDB
	c               *cron.Cron
	reactService    gRPCService.ReactionServiceClient
	regionSupported *map[string]Coordinate

	gRPCService.UnimplementedWeatherServiceServer
}

func NewWeatherService() (*WeatherService, error) {
	scheduler := cron.New()
	scheduler.Start()
	tokenDb, errTok := db.InitTokenDb()
	weatherDb, errW := db.InitWeatherConditionDb()
	if err := cmp.Or(errTok, errW); err != nil {
		return nil, err
	}
	weatherService := &WeatherService{
		tokenDb:         tokenDb,
		weatherDb:       weatherDb,
		c:               scheduler,
		reactService:    nil,
		regionSupported: RegionSupported(),
	}
	weatherService.c.AddFunc("@hourly", weatherService.checkDayCondition)
	weatherService.c.AddFunc("@hourly", weatherService.checkTemperature)
	weatherService.c.AddFunc("@hourly", weatherService.checkRain)
	weatherService.c.AddFunc("@hourly", weatherService.checkSnow)
	return weatherService, nil
}

func (weather *WeatherService) InitReactClient(conn *grpc.ClientConn) {
	weather.reactService = gRPCService.NewReactionServiceClient(conn)
}

func (weather *WeatherService) createNewWeatherInfo(
	resp *WeatherAPIResponseBody,
	userID uint,
	actionID int,
	actionType models.WeatherActionType,
	temperature float64,
	rain float64,
	snow float64,
) error {
	_, err := weather.weatherDb.InsertNewWeatherCondition(&models.WeatherCondition{
		ActionID:           uint(actionID),
		UserID:             userID,
		ActionType:         actionType,
		Activated:          true,
		Temperature:        temperature,
		TemperatureMetrics: "°C",
		Timezone:           resp.Timezone,
		Latitude:           resp.Latitude,
		Longitude:          resp.Longitude,
		Rain:               rain,
		SnowFall:           snow,
		IsDay:              resp.Current.IsDay,
	})
	return err
}

func (weather *WeatherService) NewTemperatureTrigger(ctx context.Context, req *gRPCService.TempTriggerReq) (*gRPCService.TempTriggerReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "WeatherService")
	if err != nil {
		return nil, err
	}

	if coord, ok := (*weather.regionSupported)[req.Region]; ok {
		resp, err := GetCurrentWeather(&WeatherConfig{
			Latitude:  coord.Latitude,
			Longitude: coord.Longitude,
			Current:   "temperature_2m",
			Timezone:  req.Timezone,
		})
		if err != nil {
			log.Println("Could not fetch weather data: ", err)
			return nil, err
		}
		err = weather.createNewWeatherInfo(resp, userID, int(req.ActionId), models.TemperatureExceed, float64(req.Temperature), resp.Current.Rain, resp.Current.SnowFall)
		if err != nil {
			return nil, err
		}
		log.Println("Temperature is looked at")
	} else {
		return nil, status.Errorf(codes.NotFound, "Region: %v is not supported", req.Region)
	}
	return req, nil
}

func (weather *WeatherService) NewIsDayTrigger(ctx context.Context, req *gRPCService.IsDayTriggerReq) (*gRPCService.IsDayTriggerReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "WeatherService")
	if err != nil {
		return nil, err
	}

	if coord, ok := (*weather.regionSupported)[req.Region]; ok {
		resp, err := GetCurrentWeather(&WeatherConfig{
			Latitude:  coord.Latitude,
			Longitude: coord.Longitude,
			Current:   "is_day",
			Timezone:  req.Timezone,
		})
		if err != nil {
			log.Println("Could not fetch weather data: ", err)
			return nil, err
		}
		err = weather.createNewWeatherInfo(resp, userID, int(req.ActionId), models.DayCondition, resp.Current.Temperature2m, resp.Current.Rain, resp.Current.SnowFall)
		if err != nil {
			return nil, err
		}
		log.Println("Day is looked at")
	} else {
		return nil, status.Errorf(codes.NotFound, "Region: %v is not supported", req.Region)
	}
	return req, nil
}

func (weather *WeatherService) NewRainTrigger(ctx context.Context, req *gRPCService.IsRainTriggerReq) (*gRPCService.IsRainTriggerReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "WeatherService")
	if err != nil {
		return nil, err
	}

	if coord, ok := (*weather.regionSupported)[req.Region]; ok {
		resp, err := GetCurrentWeather(&WeatherConfig{
			Latitude:  coord.Latitude,
			Longitude: coord.Longitude,
			Current:   "rain",
			Timezone:  req.Timezone,
		})
		if err != nil {
			log.Println("Could not fetch weather data: ", err)
			return nil, err
		}
		err = weather.createNewWeatherInfo(resp, userID, int(req.ActionId), models.Raining, resp.Current.Temperature2m, 0, resp.Current.SnowFall)
		if err != nil {
			return nil, err
		}
		log.Println("Rain is looked at")
	} else {
		return nil, status.Errorf(codes.NotFound, "Region: %v is not supported", req.Region)
	}
	return req, nil
}

func (weather *WeatherService) NewSnowTrigger(ctx context.Context, req *gRPCService.IsSnowTriggerReq) (*gRPCService.IsSnowTriggerReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "WeatherService")
	if err != nil {
		return nil, err
	}

	if coord, ok := (*weather.regionSupported)[req.Region]; ok {
		resp, err := GetCurrentWeather(&WeatherConfig{
			Latitude:  coord.Latitude,
			Longitude: coord.Longitude,
			Current:   "snowfall",
			Timezone:  req.Timezone,
		})
		if err != nil {
			log.Println("Could not fetch weather data: ", err)
			return nil, err
		}
		err = weather.createNewWeatherInfo(resp, userID, int(req.ActionId), models.Snowing, resp.Current.Temperature2m, resp.Current.Rain, 0)
		if err != nil {
			return nil, err
		}
		log.Println("Snow is looked at")
	} else {
		return nil, status.Errorf(codes.NotFound, "Region: %v is not supported", req.Region)
	}
	return req, nil
}

func (weather *WeatherService) SetActivate(ctx context.Context, req *gRPCService.SetActivateWeather) (*gRPCService.SetActivateWeather, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "weather")
	if err != nil {
		return nil, err
	}
	_, err = weather.weatherDb.SetActivateByActionID(req.Activated, userID, uint(req.ActionId))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (weather *WeatherService) DeleteAction(ctx context.Context, req *gRPCService.DeleteWeatherActionReq) (*gRPCService.DeleteWeatherActionReq, error) {
	userID, err := grpcutils.GetUserIdFromContext(ctx, "weather")
	if err != nil {
		return nil, err
	}
	return req, weather.weatherDb.DeleteByActionID(userID, uint(req.ActionId))
}
