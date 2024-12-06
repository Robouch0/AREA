//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-oauth
//

package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type OAuthRequest struct {
	Service string `json:"service"`
	Code    string `json:"code"`
}

type OAuthAccess struct {
	Token     string `json:"access_token"`
	Scope     string `json:"scope"`
	Tokentype string `json:"token_type"`
}

func createOAuthURLS() map[string][]string {
	m := make(map[string][]string)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m["github"] = []string{
		fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user:email", os.Getenv("GITHUB_ID"), "http://localhost:3100/"),
		"https://github.com/login/oauth/access_token",
		"https://api.github.com/user/emails",
	}

	return m
}

func OAuthRoutes() chi.Router {
	OAuthRouter := chi.NewRouter()
	OAuthURL := createOAuthURLS()

	OAuthRouter.Get("/{service}", func(w http.ResponseWriter, r *http.Request) {
		OAuthservice := chi.URLParam(r, "service")

		if url, ok := OAuthURL[OAuthservice]; ok {
			w.Write([]byte(url[0]))
		}
	})

	OAuthRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		OAuthCode := new(OAuthRequest)
		err := json.NewDecoder(r.Body).Decode(&OAuthCode)

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}
		if url, ok := OAuthURL[OAuthCode.Service]; ok {
			body := fmt.Sprintf(`{ "client_id" : "%s", "client_secret" : "%s", "code" : "%s" }`,
				os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
				os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
				OAuthCode.Code)
			request, err := http.Post(url[1], "application/json", bytes.NewBuffer([]byte(body)))

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			defer request.Body.Close()
			Tknbody, _ := io.ReadAll(request.Body)
			TokenFidx := strings.TrimLeft(string(Tknbody), "access_token=")
			TokenSidx := strings.Index(TokenFidx, "&scope")

			client := &http.Client{}
			req, _ := http.NewRequest("GET", url[2], nil)
			req.Header = http.Header{
				"Accept": {"application/vnd.github+json"},
				"Authorization": {fmt.Sprintf("Bearer %s", TokenFidx[:TokenSidx])},
				"X-GitHub-Api-Version" : {"2022-11-28"},
			}
			res, err := client.Do(req)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			defer res.Body.Close()
			UsrInf, _ := io.ReadAll(res.Body)
			w.Write([]byte(string(UsrInf)))
		}

	})
	return OAuthRouter
}
/*
		curl --request GET \
		--url "" \
		--header "" \
		--header "Authorization: Bearer USER_ACCESS_TOKEN" \
		--header ": "
*/
