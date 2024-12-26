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
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Swaggoooo
func handleWebhookPayload(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service := chi.URLParam(r, "service")
		action_id := chi.URLParam(r, "action_id")

		if _, err := strconv.Atoi(action_id); err != nil {
			// Error http
			return
		}
		if cli, ok := gateway.Clients[service]; ok {
			cli.ListServiceStatus()
		}
	}
}

func WebHookRoutes(gateway *api.ApiGateway) chi.Router {
	WebHooks := chi.NewRouter()
	db.InitTokenDb()

	WebHooks.Post("/service}/{action_id}", handleWebhookPayload(gateway))
	return WebHooks
}
