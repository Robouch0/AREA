//
// EPITECH PROJECT, 2024
// AREA
// File description:
// createRoutes
//

package controllers

import (
	api "area/api"
	gRPCapi "area/gRPC/api"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func sendToService(cliService gRPCapi.ClientService, body map[string]any) (string, error) {
	msg, err := cliService.SendAction(body)
	if err != nil {
		return "", err
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
			json.Unmarshal(b, &body)

			msg, err := sendToService(service, body)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			// msg, err := sendToService(gateway.Clients["react"], body)
			// if err != nil {
			// 	w.WriteHeader(401)
			// 	w.Write([]byte(err.Error()))
			// 	return
			// }
			w.Write([]byte(msg))
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
