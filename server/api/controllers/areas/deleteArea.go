//
// EPITECH PROJECT, 2025
// AREA
// File description:
// deleteArea
//

package areas

import (
	api "area/api"
	conv_utils "area/utils/convUtils"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
)

type areaDeleteRequest struct {
	AreaID uint32 `json:"area_id"`
}

// Area godoc
// @Summary      Delete an area
// @Description  Delete user's area
// @Security ApiKeyAuth
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param 		 area body	areaDeleteRequest	true 	"Informations about the deletion of an area"
// @Success      200  {object}  serviceinterface.SetActivatedResponseStatus
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /area [delete]
func DeleteArea(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call the rpc function for delete of react service
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		areaReq, err := conv_utils.IoReaderToStruct[areaDeleteRequest](&r.Body)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		log.Printf("Deleting area (%v)\n", areaReq.AreaID)
		resp, err := gateway.Clients["react"].DeleteArea(uint(areaReq.AreaID), userID)
		if err != nil {
			log.Println("Error in reaction service: ", err)
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		b, err := json.Marshal(&resp)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		w.WriteHeader(200)
		w.Write(b)
	}
}
