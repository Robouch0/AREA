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

func GetAccessToken(OAuthCode *OAuthRequest, url []string) (resp *http.Response, err error) {
	body := fmt.Sprintf(`{ "client_id" : "%s", "client_secret" : "%s", "code" : "%s" }`,
		os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
		os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
		OAuthCode.Code)

	request, err := http.Post(url[1], "application/json", bytes.NewBuffer([]byte(body)))

	if err != nil {
		return nil, err
	}
	return request, nil
}

func GetUserEmail(url []string, TokenStr string, idx int) (resp *http.Response, err error) {
	if idx == -1 {
		return nil, err
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", url[2], nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", TokenStr[:idx]))
	result, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	return result, nil
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
			request, err := GetAccessToken(OAuthCode, url)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			defer request.Body.Close()
			Tknbody, _ := io.ReadAll(request.Body)
			TokenFidx := strings.TrimLeft(string(Tknbody), "access_token=")
			TokenSidx := strings.Index(TokenFidx, "&scope")
			result, err := GetUserEmail(url, TokenFidx, TokenSidx)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			defer result.Body.Close()
		}

	})
	return OAuthRouter
}
