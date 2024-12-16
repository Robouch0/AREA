//
// EPITECH PROJECT, 2024
// AREA
// File description:
// createRoutes
//

package controllers

import (
	api "area/api"
	gRPCapi "area/gRPC/api/serviceInterface"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	DEFAULT_ACTION_ID int = -1
)

func sendToService(cliService gRPCapi.ClientService, body map[string]any, actionID int) (*gRPCapi.ActionResponseStatus, error) {
	msg, err := cliService.SendAction(body, actionID)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

// Area godoc
// @Summary      Create a new Area
// @Description  Register a new Area in the application
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param 		 service path	string	true 	"Service Name"
// @Param 		 area body	models.AreaScenario	true 	"Full body of an Area Scenario"
// @Success      200  {object}  gRPCapi.ActionResponseStatus
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /create/{service} [get]
func CreateRoute(gateway *api.ApiGateway) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		serviceParam := chi.URLParam(r, "service")

		if service, ok := gateway.Clients[serviceParam]; ok {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				log.Println(err)
				return
			}

			body := map[string]any{}
			err = json.Unmarshal(b, &body)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				log.Println(err)
				return
			}

			msg, err := sendToService(gateway.Clients["react"], body, DEFAULT_ACTION_ID) // Create the action if possible
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				log.Println(err)
				return
			}
			msg, err = sendToService(service, body, msg.ActionID)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				log.Println(err)
				return
			}

			res, err := json.Marshal(msg)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(res))
		} else {
			w.WriteHeader(401)
			w.Write([]byte(fmt.Sprintf("No such Service: %v", serviceParam)))
		}
	}
}
