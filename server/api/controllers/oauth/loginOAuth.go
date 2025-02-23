//
// EPITECH PROJECT, 2024
// AREA
// File description:
// loginOAuth
//

package oauth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"

	_ "area/api/controllers/log_types"
	http_utils "area/utils/httpUtils"
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
func LoginOAuth(JwtTok *jwtauth.JWTAuth, OAuthURL map[string]OAuthURLs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		OAuthCode := new(OAuthRequest)
		err := json.NewDecoder(r.Body).Decode(&OAuthCode)

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			log.Println("Cannot parse JSON OAuth Request")
			return
		}
		if oauth, ok := OAuthURL[OAuthCode.Service]; ok {
			toks, err := oauth.OAuth.GetAccessToken(OAuthCode)
			if err != nil {
				log.Println(err)
				http_utils.WriteHTTPResponseErr(&w, 500, err.Error())
				return
			}
			log.Println("Token: ", toks.AccessToken)
			if err = oauth.OAuth.HandleUserTokens(*toks, &w, JwtTok); err != nil {
				log.Println(err)
				http_utils.WriteHTTPResponseErr(&w, 500, err.Error())
				return
			}
			return
		}
		http_utils.WriteHTTPResponseErr(&w, 404, fmt.Sprintf("Service %s not found", OAuthCode.Service))
	}
}
