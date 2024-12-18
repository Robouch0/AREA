//
// EPITECH PROJECT, 2024
// AREA
// File description:
// tokens
//

package controllers

import (
	"area/db"
	"area/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TokenRequest struct {
	UserID   string `json:"user_id"`
	Provider string `json:"provider"`
}

type TokenCreateRequest struct {
	UserID   string `json:"user_id"`
	Provider string `json:"provider"`
	Token    string `json:"token"`
}

// Get tokens godoc
// @Summary      Get all the tokens from a user
// @Description  Get all the tokens from a user_id
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/{user_id} [get]
func getTokens(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		struserID := chi.URLParam(r, "user_id")
		userID, err := strconv.Atoi(struserID)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		tokens, err := tokenDb.GetUserTokens(int64(userID))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tokens)
	}
}

// Get token godoc
// @Summary      Get a token
// @Description  Get the tokens from a user_id and a provider
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/ [post]
func getToken(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		TokenReq := new(TokenRequest)
		TokenReq.UserID = chi.URLParam(r, "user_id")
		TokenReq.Provider = chi.URLParam(r, "provider")

		UserId, err := strconv.Atoi(TokenReq.UserID)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		tokens, err := tokenDb.GetToken(int64(UserId), TokenReq.Provider)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tokens)
	}
}

// create Token godoc
// @Summary      Create a token
// @Description  Create a token from a user_id and a provider
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/create/ [post]
func createTkn(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateReq := new(TokenCreateRequest)
		var NewToken models.Token
		err := json.NewDecoder(r.Body).Decode(&CreateReq)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		UserId, err := strconv.Atoi(CreateReq.UserID)
		Oldtoken, err := tokenDb.GetToken(int64(UserId), CreateReq.Provider)

		if err != nil {
			NewToken.AccessToken = CreateReq.Token
			NewToken.Provider = CreateReq.Provider
			NewToken.UserID = int64(UserId)
			token, err := tokenDb.CreateToken(&NewToken)

			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(&token)
		} else {
			Oldtoken.AccessToken = CreateReq.Token
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(&Oldtoken)
		}
	}
}

func TokenRoutes() chi.Router {
	TokenRouter := chi.NewRouter()
	tokenDb := db.GetTokenDb()

	TokenRouter.Get("/{user_id}", getTokens(tokenDb))

	TokenRouter.Get("/", getToken(tokenDb))

	TokenRouter.Post("/create/", createTkn(tokenDb))

	return TokenRouter
}
