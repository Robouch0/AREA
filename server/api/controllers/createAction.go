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

func CreateRoute(gateway *api.ApiGateway) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		serviceParam := chi.URLParam(r, "service")

		if service, ok := gateway.Clients[serviceParam]; ok {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}

			body := map[string]any{}
			err = json.Unmarshal(b, &body)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}

			msg, err := sendToService(gateway.Clients["react"], body, DEFAULT_ACTION_ID) // Create the action if possible
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			msg, err = sendToService(service, body, msg.ActionID)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
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

/*
{
    "user_id": int,
    "action": {
        "service": string,
        "microservice": string,
        "ingredients": jsonb,
    },
	"reaction": {
        "service": string,
        "microservice": string,
        "ingredients": jsonb,
    },
}
*/
