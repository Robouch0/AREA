//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-oauth
//

package controllers

import (
	"area/api/middleware"
	"area/db"
	"area/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
)

type OAuthRequest struct {
	Service string `json:"service"`
	Code    string `json:"code"`
}

type OAuthInfos struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
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

func GetUserEmail(url []string, TokenStr string, idx int, w http.ResponseWriter) (resp *http.Response, err error) {
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

func createUser(w http.ResponseWriter, userlist []OAuthInfos, JwtTok *jwtauth.JWTAuth) (*models.User, error) {
	userDb := db.GetUserDb()
	us := new(models.User)

	err := userDb.Db.NewSelect().
		Model(us).
		Where("email = ?", userlist[0].Email).
		Scan(context.Background())
	if err != nil {
		var newUser models.User
		newUser.Email = userlist[0].Email
		_, err := userDb.CreateUser(&newUser)

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return nil, err
		}
	}
	w.WriteHeader(200)
	w.Write([]byte(middleware.CreateToken(JwtTok, us.ID)))
	return us, nil
}

func createToken(w http.ResponseWriter, user *models.User, AccessToken string, Service string) {
	tokenDb := db.GetTokenDb()
	tkn := new(models.Token)
	err := tokenDb.Db.NewSelect().
		Model(tkn).
		Where("user_id = ? AND provider = ? ", user.ID, Service).
		Scan(context.Background())

	if err != nil {
		var newToken models.Token
		newToken.Provider = Service
		newToken.UserID = int64(user.ID)
		newToken.AccessToken = AccessToken
		_, err := tokenDb.CreateToken(&newToken)

		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		tkn.AccessToken = AccessToken
	}
}

func OAuthRoutes(JwtTok *jwtauth.JWTAuth) chi.Router {
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
			result, err := GetUserEmail(url, TokenFidx, TokenSidx, w)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			userlist := []OAuthInfos{}
			err = json.NewDecoder(result.Body).Decode(&userlist)

			if err != nil {
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			user, err := createUser(w, userlist, JwtTok)
			if err == nil {
				createToken(w, user, TokenFidx[:TokenSidx], OAuthCode.Service)
			}
		}

	})
	return OAuthRouter
}
