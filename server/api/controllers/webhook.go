//
// EPITECH PROJECT, 2024
// AREA
// File description:
// webhook
//

package controllers

import (
	"area/api"
	"area/db"
	"area/utils"
	http_utils "area/utils/httpUtils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Swaggoooo
func handleWebhookPayload(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := chi.URLParam(r, "service")
		microservice := chi.URLParam(r, "microservice")
		action_id := chi.URLParam(r, "action_id")

		actionId, err := strconv.Atoi(action_id)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Incorrect format ")
			return
		}
		payload, err := utils.IoReaderToMap(&r.Body)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Invalid json payload")
			return
		}
		if cli, ok := gateway.Clients[service]; ok {
			cli.TriggerWebhook(payload, microservice, actionId)
		}
		w.WriteHeader(200)
		w.Write([]byte("Done"))
	}
}

func WebHookRoutes(gateway *api.ApiGateway) chi.Router {
	WebHooks := chi.NewRouter()
	db.InitTokenDb()

	WebHooks.Post("/service}/{microservice}/{action_id}", handleWebhookPayload(gateway))
	return WebHooks
}
