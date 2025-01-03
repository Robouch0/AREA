//
// EPITECH PROJECT, 2024
// AREA
// File description:
// areas
//

package controllers

import (
	api "area/api"
	"area/db"
	IServ "area/gRPC/api/serviceInterface"
	"area/models"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
	"slices"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userArea struct {
	ID        uint                  `json:"id,omitempty"`
	Action    IServ.ServiceStatus   `json:"action,omitempty"`
	Reactions []IServ.ServiceStatus `json:"reactions,omitempty"`
	Activated bool                  `json:"activated,omitempty"`
}

func getCorrectMicroservice(gateway *api.ApiGateway, serviceStatus *IServ.ServiceStatus, microservice string) (*IServ.MicroserviceDescriptor, error) {
	idx := slices.IndexFunc(serviceStatus.Microservices, func(micro IServ.MicroserviceDescriptor) bool {
		return micro.RefName == microservice
	})
	if idx == -1 {
		return nil, status.Errorf(codes.NotFound, "Microservice %v not found", microservice)
	}
	return &serviceStatus.Microservices[idx], nil
}

func getActionArea(gateway *api.ApiGateway, action *models.Action) (IServ.ServiceStatus, error) {
	serviceStatus, err := gateway.Clients[action.Service].ListServiceStatus()
	if err != nil {
		return IServ.ServiceStatus{}, err
	}

	micro, err := getCorrectMicroservice(gateway, serviceStatus, action.Microservice)
	if err != nil {
		return IServ.ServiceStatus{}, err
	}

	for k, v := range (*micro).Ingredients {
		value, ok := action.Ingredients[k]
		if !ok {
			return IServ.ServiceStatus{}, status.Errorf(codes.DataLoss, "Invalid ingredient %v", k)
		}
		v.Value = value
		(*micro).Ingredients[k] = v
	}
	return IServ.ServiceStatus{
		Name:    serviceStatus.Name,
		RefName: serviceStatus.RefName,
		Microservices: []IServ.MicroserviceDescriptor{
			*micro,
		},
	}, err
}

func getReactionsArea(gateway *api.ApiGateway, reactions []*models.Reactions) ([]IServ.ServiceStatus, error) {
	var services []IServ.ServiceStatus
	for _, react := range reactions {
		serviceStatus, err := gateway.Clients[react.Reaction.Service].ListServiceStatus()
		if err != nil || serviceStatus == nil {
			continue
		}

		micro, err := getCorrectMicroservice(gateway, serviceStatus, react.Reaction.Microservice)
		if err != nil {
			return []IServ.ServiceStatus{}, err
		}

		for k, v := range (*micro).Ingredients {
			value, ok := react.Reaction.Ingredients[k]
			if !ok {
				return []IServ.ServiceStatus{}, status.Errorf(codes.DataLoss, "Invalid ingredient %v", k)
			}
			v.Value = value
			(*micro).Ingredients[k] = v
		}
		services = append(services, IServ.ServiceStatus{
			Name:    serviceStatus.Name,
			RefName: serviceStatus.RefName,
			Microservices: []IServ.MicroserviceDescriptor{
				*micro,
			},
		})
	}
	return services, nil
}

func formatUsersArea(gateway *api.ApiGateway, areas *[]models.Area) ([]userArea, error) {
	var allAreas []userArea
	for _, area := range *areas {
		action, errAct := getActionArea(gateway, area.Action.Action)
		if errAct != nil {
			return nil, errAct
		}
		reactions, errReact := getReactionsArea(gateway, area.Reactions)
		if errReact != nil {
			return nil, errReact
		}

		allAreas = append(allAreas, userArea{
			ID:        area.ID,
			Action:    action,
			Reactions: reactions,
			Activated: area.Activated,
		})
	}
	return allAreas, nil
}

// Area godoc
// @Summary      List User's area
// @Description  List all user's area
// @Tags         Area
// @Accept       json
// @Produce      json
// @Success      200  {object}  []userArea
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /area/list [get]
func GetUserAreas(gateway *api.ApiGateway, areaDB *db.AreaDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		areas, err := areaDB.GetFullAreaByUserID(userID)
		if err != nil {
			log.Println("Error while loading all the user's areas", err)
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		allAreas, err := formatUsersArea(gateway, areas)
		if err != nil {
			log.Println("Error while formatting all of nthe user's areas", err)
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		b, err := json.Marshal(allAreas)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}
