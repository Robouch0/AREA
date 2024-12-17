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
	"area/utils"
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

	Primary    bool   `json:"primary,omitempty"`
	Verified   bool   `json:"verified,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type OAuthURLs struct {
	RedirectURL     string `json:"redirect_uri"`
	AccessTokenURL  string `json:"access_token_uri"`
	EmailRequestURL string `json:"email_req_uri"`
}

func createOAuthRedirect(consentURL, provider, redirectURI string) string {
	return fmt.Sprintf(consentURL, os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(provider))), redirectURI)
}

func createOAuthURLS() map[string]OAuthURLs {
	oauthUrls := make(map[string]OAuthURLs)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Mettre manuellement tous les scopes
	oauthUrls["github"] = OAuthURLs{
		RedirectURL:     "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user:email,repo,workflow", // Mettre plus de droits précis
		AccessTokenURL:  "https://github.com/login/oauth/access_token",
		EmailRequestURL: "https://api.github.com/user/emails",
	}
	oauthUrls["google"] = OAuthURLs{
		RedirectURL:     "https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&scope=%s&response_type=code",
		AccessTokenURL:  "https://oauth2.googleapis.com/token", // https://accounts.google.com/o/oauth2/token
		EmailRequestURL: "https://www.googleapis.com/oauth2/v1/userinfo",
	}
	return oauthUrls
}

func GetAccessToken(OAuthCode *OAuthRequest, accessTokenURI string) (resp *http.Response, err error) {
	body := fmt.Sprintf(`{ "client_id" : "%s", "client_secret" : "%s", "code" : "%s" }`,
		os.Getenv(fmt.Sprintf("%s_ID", strings.ToUpper(OAuthCode.Service))),
		os.Getenv(fmt.Sprintf("%s_SECRET", strings.ToUpper(OAuthCode.Service))),
		OAuthCode.Code)

	request, err := http.Post(accessTokenURI, "application/json", bytes.NewBuffer([]byte(body)))

	if err != nil {
		return nil, err
	}
	return request, nil
}

func GetUserEmail(emailRequestURL string, TokenStr string, idx int, w http.ResponseWriter) (resp *http.Response, err error) {
	if idx == -1 {
		return nil, err
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", emailRequestURL, nil)
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
			utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return nil, err
		}
		us = &newUser
	}
	b, err := json.Marshal(UserLogInfos{Token: middleware.CreateToken(JwtTok, us.ID), UserID: us.ID})
	if err != nil {
		utils.WriteHTTPResponseErr(&w, 401, err.Error())
		return nil, err
	}
	w.WriteHeader(200)
	w.Write(b)
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
		OAuthRedirectURI := r.URL.Query().Get("redirect_uri") // Rajouter dans le swagger les urls params

		if serviceUrls, ok := OAuthURLs[OAuthservice]; ok {
			url := createOAuthRedirect(serviceUrls.RedirectURL, OAuthservice, OAuthRedirectURI)
			log.Println("URL: ", url)
			if url != "" {
				w.WriteHeader(200)
				w.Write([]byte(url))
			} else {
				utils.WriteHTTPResponseErr(&w, 400, "Service does not exist")
			}
		}
	}
}

// Sign-up OAuth godoc
// @Summary      Create account with oauth
// @Description  Create account with code from redirect url
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      200  {object}  jwtauth.JWTAuth
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /oauth/ [post]
func LoginOAuth(JwtTok *jwtauth.JWTAuth, OAuthURL map[string]OAuthURLs) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		OAuthCode := new(OAuthRequest)
		err := json.NewDecoder(r.Body).Decode(&OAuthCode)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		if urls, ok := OAuthURL[OAuthCode.Service]; ok {
			request, err := GetAccessToken(OAuthCode, urls.AccessTokenURL)

			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			defer request.Body.Close()
			Tknbody, _ := io.ReadAll(request.Body)
			TokenFidx := strings.TrimLeft(string(Tknbody), "access_token=")
			TokenSidx := strings.Index(TokenFidx, "&scope")
			result, err := GetUserEmail(urls.EmailRequestURL, TokenFidx, TokenSidx, w)

			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			userlist := []OAuthInfos{}
			err = json.NewDecoder(result.Body).Decode(&userlist)

			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			user, err := createUser(w, userlist, JwtTok)
			if err == nil {
				createToken(w, user, TokenFidx[:TokenSidx], OAuthCode.Service)
			}
		}
	}
}

func OAuthRoutes(JwtTok *jwtauth.JWTAuth) chi.Router {
	OAuthRouter := chi.NewRouter()
	OAuthURLs := createOAuthURLS()
	db.InitTokenDb()

	OAuthRouter.Get("/{service}", GetUrl(OAuthURLs))

	OAuthRouter.Post("/", LoginOAuth(JwtTok, OAuthURLs))
	return OAuthRouter
}
