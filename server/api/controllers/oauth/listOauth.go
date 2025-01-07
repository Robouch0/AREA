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

func ListOauth(OAuthURLs map[string]OAuthURLs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Oauths []string
		for key, _ := range OAuthURLs {
			Oauths = append(Oauths, key)
		}
		json.NewEncoder(w).Encode(Oauths)
	}
}
