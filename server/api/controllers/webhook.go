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

// WebHook Enpoint godoc
// @Summary      WebHook Enpoint
// @Description  WebHook Enpoint for the remote services payloads
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param 		 service      path	string	true 	"Service Name"
// @Param 		 microservice path	string	true 	"Microservice Name"
// @Param 		 action_id    path	string	true 	"Action ID for the reaction service"
// @Success      200  {object}  string
// @Failure      401  {object}  error
// @Router       /webhook/{service}/{microservice}/{action_id} [post]
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
			w.WriteHeader(204)
			return
		}
		http_utils.WriteHTTPResponseErr(&w, 401, "Invalid payload sent")
	}
}

// Créer une db google gmail

func WebHookRoutes(gateway *api.ApiGateway) chi.Router {
	WebHooks := chi.NewRouter()
	db.InitTokenDb()

	WebHooks.Post("/{service}/{microservice}/{action_id}", handleWebhookPayload(gateway))
	return WebHooks
}
