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

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	Clientdata := ClientInfo{
		Host: r.RemoteAddr,
		Time: time.Now().Unix(),
	}
	fmt.Printf(r.Header.Get("X-Real-Ip"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Clientdata)
}
