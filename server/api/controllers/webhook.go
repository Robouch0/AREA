//
// EPITECH PROJECT, 2024
// AREA
// File description:
// webhook
//

package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func WebHookRoute(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	action_id := chi.URLParam(r, "action_id")

	log.Println("HERE: ", service, action_id)
	io.Copy(os.Stdout, r.Body)
}
