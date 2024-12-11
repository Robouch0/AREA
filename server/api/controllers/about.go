//
// EPITECH PROJECT, 2024
// AREA
// File description:
// about
//

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ClientInfo struct {
	Host string
	Time int64
}

// AboutRoute godoc
// @Summary      List of handled services
// @Description  json giving the list of handled action-reaction services
// @Tags         Utils
// @Accept       json
// @Produce      json
// @Router       /about.json [get]
func AboutRoute(w http.ResponseWriter, r *http.Request) {
	Clientdata := ClientInfo{
		Host: r.RemoteAddr,
		Time: time.Now().Unix(),
	}
	fmt.Printf(r.Header.Get("X-Real-Ip"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Clientdata)
}
