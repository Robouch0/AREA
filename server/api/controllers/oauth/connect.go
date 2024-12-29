//
// EPITECH PROJECT, 2024
// AREA
// File description:
// connect
//

package oauth

import (
	"area/db"
	grpcutils "area/utils/grpcUtils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Sign-up OAuth godoc
// @Summary      Create account with oauth
// @Description  Create account with code from redirect url
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  log_types.UserLogInfos
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Failure      404  {object}  error
// @Router       /oauth/ [post]

func Connect(OAuthURL map[string]OAuthURLs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Invalid claims in jwt tokens.")
			log.Println("Invalid claims in jwt tokens", err)
			return
		}
		OAuthCode := new(OAuthRequest)
		err = json.NewDecoder(r.Body).Decode(&OAuthCode)

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			log.Println("Cannot parse JSON OAuth Request")
			return
		}
		userDb := db.GetUserDb()
		user, err := userDb.GetUserByID(int(userID))

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			log.Println("Cannot get user by id")
			return
		}

		if oauth, ok := OAuthURL[OAuthCode.Service]; ok {
			toks, err := oauth.OAuth.GetAccessToken(OAuthCode)
			if err != nil {
				log.Println(err)
				http_utils.WriteHTTPResponseErr(&w, 500, err.Error())
				return
			}

			CreateToken(w, user, toks.AccessToken, OAuthCode.Service)
			return
		}
		http_utils.WriteHTTPResponseErr(&w, 404, fmt.Sprintf("Service %s not found", OAuthCode.Service))
	}
}
