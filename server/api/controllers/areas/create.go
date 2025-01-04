//
// EPITECH PROJECT, 2024
// AREA
// File description:
// createRoutes
//

package areas

import (
	api "area/api"
	gRPCapi "area/gRPC/api/serviceInterface"
	"area/models"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	DEFAULT_ACTION_ID int = -1
)

func sendToService(cliService gRPCapi.ClientService, scenario models.AreaScenario, actionID int, userID int) (*gRPCapi.ActionResponseStatus, error) {
	msg, err := cliService.SendAction(scenario, actionID, userID)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

// Area godoc
// @Summary      Create a new Area
// @Description  Register a new Area in the application
// @Tags         Area
// @Accept       json
// @Produce      json
// @Param 		 service path	string	true 	"Service Name"
// @Param 		 area body	models.AreaScenario	true 	"Full body of an Area Scenario"
// @Success      200  {object}  gRPCapi.ActionResponseStatus
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /area/create/{service} [post]
func CreateArea(gateway *api.ApiGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceParam := chi.URLParam(r, "service")
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Invalid claims in jwt tokens.")
			log.Println("Invalid claims in jwt tokens", err)
			return
		}
		if service, ok := gateway.Clients[serviceParam]; ok {
			var scenario models.AreaScenario

			err := json.NewDecoder(r.Body).Decode(&scenario)
			if err != nil {
				http_utils.WriteHTTPResponseErr(&w, 401, "Incorrect body is sent.")
				log.Printf("Json Error: %v\n", err)
				return
			}

			msg, err := sendToService(gateway.Clients["react"], scenario, DEFAULT_ACTION_ID, int(userID))
			if err != nil {
				http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
				log.Println("ReactionService error: ", err)
				return
			}
			msg, err = sendToService(service, scenario, msg.ActionID, int(userID))
			if err != nil {
				http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
				log.Println("ActionService error: ", err)
				return
			}

			res, err := json.Marshal(msg)
			if err != nil {
				http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
				log.Println(err)
				return
			}
			w.Write([]byte(res))
		} else {
			http_utils.WriteHTTPResponseErr(&w, 401, fmt.Sprintf("No such Service: %v", serviceParam))
		}
	}
}
