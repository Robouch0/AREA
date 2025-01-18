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
	serviceinterface "area/gRPC/api/serviceInterface"
	conv_utils "area/utils/convUtils"
	http_utils "area/utils/httpUtils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Webhook Enpoint godoc
// @Summary      Webhook Endpoint
// @Description  Webhook Endpoint for the remote services payloads
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
		payload, err := conv_utils.IoReaderToMap(&r.Body)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Invalid json payload")
			return
		}
		if cli, ok := gateway.Clients[service]; ok {
			_, err := cli.TriggerWebhook(&serviceinterface.WebhookInfos{Payload: payload, Header: r.Header}, microservice, actionId)
			if err != nil {
				log.Println(err)
				http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
				return
			}
			w.WriteHeader(204)
			return
		}
		http_utils.WriteHTTPResponseErr(&w, 401, "Invalid payload sent")
	}
}

func WebHookRoutes(gateway *api.ApiGateway) chi.Router {
	WebHooks := chi.NewRouter()
	db.InitTokenDb()

	WebHooks.Post("/{service}/{microservice}/{action_id}", handleWebhookPayload(gateway))
	return WebHooks
}
