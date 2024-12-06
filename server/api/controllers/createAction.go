//
// EPITECH PROJECT, 2024
// AREA
// File description:
// createRoutes
//

package controllers

import (
	api "area/api"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
{
	"scenario":
}

// Faire les types pour la gateway et les envois aux services
// Avec des maps
*/

func CreateRoute(gateway *api.ApiGateway) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		serviceParam := chi.URLParam(r, "service")

		if service, ok := gateway.Clients[serviceParam]; ok {
			var b []byte

			_, err := r.Body.Read(b)
			if err != nil {
				if err != nil {
					w.WriteHeader(401)
					w.Write([]byte(err.Error()))
					return
				}
			}
			msg, err := service.SendAction(b)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write([]byte(msg))

		} else {
			w.WriteHeader(401)
			w.Write([]byte(fmt.Sprintf("No such Service: %v", serviceParam)))
		}
	}
}
