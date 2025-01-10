//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// listOauth
//

package oauth

import (
	"encoding/json"
	"net/http"
)

type Oauthlist struct {
	Services []string `json:"services"`
}

// Sign-up OAuth godoc
// @Summary      List Oauth
// @Description  List all the current Oauth handled by the server
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  []string
// @Router       /oauth/oauthlist [get]
func ListOauth(OAuthURLs map[string]OAuthURLs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Oauths Oauthlist
		for key, _ := range OAuthURLs {
			Oauths.Services = append(Oauths.Services, key)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(Oauths)
	}
}
