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
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type OAuthRequest struct {
	Service string `json:"service"`
	Code    string `json:"code"`
}

func createOAuthSignURL() map[string]string {
	m := make(map[string]string)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m["github"] = fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", os.Getenv("GITHUB_ID"), "http://localhost:3100/")

	return m
}

func createOAuthTokenURL() map[string]string {
	m := make(map[string]string)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m["github"] = "https://github.com/login/oauth/access_token"

	return m
}

func OAuthRoutes() chi.Router {
	OAuthRouter := chi.NewRouter()
	OAuthSignURL := createOAuthSignURL()
	OAuthTokenURL := createOAuthTokenURL()

	OAuthRouter.Get("/{service}", func(w http.ResponseWriter, r *http.Request) {
		OAuthservice := chi.URLParam(r, "service")

		if url, ok := OAuthSignURL[OAuthservice]; ok {
			w.Write([]byte(url))
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
		if url, ok := OAuthTokenURL[OAuthCode.Service]; ok {
			body := fmt.Sprintf(`{ "client_id" : "%s", "client_secret" : "%s", "code" : "%s" }`, os.Getenv("GITHUB_ID"), os.Getenv("GITHUB_SECRET"), OAuthCode.Code)
			request , err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}

			request.Header.Add("Accept", "application/json")

			client := &http.Client{}
			res, err := client.Do(request)
			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}

			w.Write([]byte(body))
			res.Body.Close()
		}

	})
	return OAuthRouter
}
