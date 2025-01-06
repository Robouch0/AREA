//
// EPITECH PROJECT, 2025
// AREA
// File description:
// activateArea
//

package areas

import (
	api "area/api"
	"area/utils"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
)

type areaActivateRequest struct {
	AreaID    uint32 `json:"area_id"`
	Activated bool   `json:"activated"`
}

// Area godoc
// @Summary      Activate/Deactivate an area
// @Description  Activate/Deactivate user's area
// @Security ApiKeyAuth
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param 		 area body	areaActivateRequest	true 	"Informations about the activation of an area"
// @Success      200  {object}  serviceinterface.SetActivatedResponseStatus
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /area/activate [put]
func ActivateArea(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		areaReq, err := utils.IoReaderToStruct[areaActivateRequest](&r.Body)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		log.Printf("Setting area (%v) activated to %v\n", areaReq.AreaID, areaReq.Activated)
		resp, err := gateway.Clients["react"].SetActivate("", uint(areaReq.AreaID), int(userID), areaReq.Activated)
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
