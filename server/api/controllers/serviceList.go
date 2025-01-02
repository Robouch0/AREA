//
// EPITECH PROJECT, 2024
// AREA
// File description:
// serviceList
//

package controllers

import (
	api "area/api"
	IServ "area/gRPC/api/serviceInterface"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"net/http"
)

// Area godoc
// @Summary      List User's area
// @Description  List all user's area
// @Tags         Area
// @Accept       json
// @Produce      json
// @Success      200  {object}  []IServ.ServiceStatus
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /create/list [get]
func ListService(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var services []IServ.ServiceStatus

		for _, cliServ := range gateway.Clients {
			servInfo, err := cliServ.ListServiceStatus()
			if err != nil || servInfo == nil {
				continue
			}
			services = append(services, *servInfo)
		}
		b, err := json.Marshal(services)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}
