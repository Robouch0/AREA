//
// EPITECH PROJECT, 2024
// AREA
// File description:
// urls
//

package oauth

import (
	"area/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func createOAuthRedirect(consentURL, provider, redirectURI string) string {
	return fmt.Sprintf(consentURL, os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(provider))), redirectURI)
}

func CreateOAuthURLS() map[string]OAuthURLs {
	oauthUrls := make(map[string]OAuthURLs)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oauthUrls["github"] = OAuthURLs{
		RedirectURL: "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user:email,repo,workflow", // Mettre plus de droits précis

		OAuth: &GithubOAuth{
			AccessTokenURL:  "https://github.com/login/oauth/access_token",
			EmailRequestURL: "https://api.github.com/user/emails",
		},
	}
	oauthUrls["google"] = OAuthURLs{ // Look again which scope is good for google for user
		RedirectURL:     "https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&scope=https://www.googleapis.com/auth/gmail.modify&response_type=code",
		AccessTokenURL:  "https://oauth2.googleapis.com/token", // https://accounts.google.com/o/oauth2/token
		AccessTokenBody: `{ "client_id" : "%s", "client_secret" : "%s", "code" : "%s", redirect_uri: "%s", grant_type: "authorization_code" }`,
		EmailRequestURL: "hthttps://www.googleapis.com/plus/v1/people/me",
	}
	return oauthUrls
}

// Sign-up OAuth godoc
// @Summary      get Oauth url by service
// @Description  get the oauth redirect url for a service
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      400  {object}  error
// @Router       /oauth/{service} [get]
func GetUrl(OAuthURLs map[string]OAuthURLs) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		OAuthservice := chi.URLParam(r, "service")
		log.Println(OAuthservice)
		OAuthRedirectURI := r.URL.Query().Get("redirect_uri") // Rajouter dans le swagger les urls params

		log.Println("Service: ", OAuthservice)
		if serviceUrls, ok := OAuthURLs[OAuthservice]; ok {
			url := createOAuthRedirect(serviceUrls.RedirectURL, OAuthservice, OAuthRedirectURI)
			if url != "" {
				w.WriteHeader(200)
				w.Write([]byte(url))
			} else {
				utils.WriteHTTPResponseErr(&w, 400, "Service does not exist")
			}
		}
	}
}
