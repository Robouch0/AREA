//
// EPITECH PROJECT, 2024
// AREA
// File description:
// ping
//

package controllers

import (
	"fmt"
	"net/http"
)

// Ping godoc
// @Summary      prints pong
// @Description  pong
// @Tags         ping
// @Produce      json
// @Success      200  "pong"
// @Router       /ping [get]
func PingRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Pong")))
}
