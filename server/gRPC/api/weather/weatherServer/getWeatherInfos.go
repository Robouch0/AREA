//
// EPITECH PROJECT, 2024
// AREA
// File description:
// getWeatherInfos
//

package weather_server

import (
	"area/utils"
	http_utils "area/utils/httpUtils"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	weatherAPIURL = "https://api.open-meteo.com/v1/forecast"
)

func RegionSupported() *map[string]Coordinate {
	return &map[string]Coordinate{
		"France":      {46.0, 2.0},
		"Europe":      {48.691, 9.1406},
		"South Korea": {36.50, 127.75},
		"Sweden":      {62, 15},
		"Finland":     {64, 26},
	}
}

func GetCurrentWeather(config *WeatherConfig) (*WeatherAPIResponseBody, error) {
	req, err := http.NewRequest("GET", weatherAPIURL, nil)
	if err != nil {
		log.Println("Request creation error: ", err)
		return nil, status.Errorf(codes.Internal, "Could not create the request: %v", err)
	}
	q := req.URL.Query()
	q.Add("latitude", strconv.FormatFloat(config.Latitude, 'f', -1, 64))
	q.Add("longitude", strconv.FormatFloat(config.Longitude, 'f', -1, 64))
	q.Add("timezone", config.Timezone)
	q.Add("current", config.Current)

	req.URL.RawQuery = q.Encode()

	resp, err := http_utils.SendHttpRequest(req, 200)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	weatherRes, err := utils.IoReaderToStruct[WeatherAPIResponseBody](&resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return weatherRes, nil
}
