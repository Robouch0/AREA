//
// EPITECH PROJECT, 2024
// AREA
// File description:
// about
//

package controllers

import (
	"area/api"
	serviceinterface "area/gRPC/api/serviceInterface"
	"encoding/json"
	"net/http"
	"time"
)

// Remote client informations
type clientInfos struct {
	Host string
}

// Microservice general informations
type microservice struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Service informations (name and actions/reactions)
type serverService struct {
	Name      string         `json:"name"`
	Actions   []microservice `json:"actions"`
	Reactions []microservice `json:"reactions"`
}

// Whole server informations
type serverInfos struct {
	CurrentTime int64           `json:"current_time"`
	Services    []serverService `json:"services"`
}

// All the about infos requested
type AboutInfos struct {
	Client clientInfos `json:"client"`
	Server serverInfos `json:"server"`
}

///////

func getCurrentInfos(r *http.Request, about *AboutInfos) {
	about.Client.Host = r.RemoteAddr
	about.Server.CurrentTime = time.Now().Unix()
}

func formatServiceStatus(infos *serviceinterface.ServiceStatus) serverService {
	serv := serverService{
		Name:      infos.Name,
		Actions:   []microservice{},
		Reactions: []microservice{},
	}
	for _, micro := range infos.Microservices {
		if micro.Type == "action" {
			serv.Actions = append(serv.Actions, microservice{
				Name:        micro.RefName,
				Description: micro.Name,
			})
		} else {
			serv.Reactions = append(serv.Reactions, microservice{
				Name:        micro.RefName,
				Description: micro.Name,
			})
		}
	}
	return serv
}

// AboutRoute godoc
// @Summary      List of handled services
// @Description  json giving the list of handled action-reaction services
// @Tags         About
// @Accept       json
// @Produce      json
// @Success      200  {object}  AboutInfos
// @Router       /about.json [get]
func AboutRoute(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		about := &AboutInfos{}
		getCurrentInfos(r, about)
		for _, cliServ := range gateway.Clients {
			servInfo, err := cliServ.ListServiceStatus()
			if err != nil || servInfo == nil {
				continue
			}
			about.Server.Services = append(about.Server.Services, formatServiceStatus(servInfo))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(about)
	}
}
