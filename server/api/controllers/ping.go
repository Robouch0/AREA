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

func PingRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Pong")))
}
