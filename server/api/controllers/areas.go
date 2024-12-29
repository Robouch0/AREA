//
// EPITECH PROJECT, 2024
// AREA
// File description:
// areas
//

package controllers

import (
	api "area/api"
	"area/db"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"net/http"
)

// Area godoc
// @Summary      List User's area
// @Description  List all user's area
// @Tags         Area
// @Accept       json
// @Produce      json
// @Success      200  {object}  gRPCapi.ActionResponseStatus
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /area/list [get]
func GetUserAreas(gateway *api.ApiGateway, areaDB *db.AreaDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		_, err = areaDB.GetAreaByUserID(userID)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
	}
}
