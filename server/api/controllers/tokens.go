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

func TokenRoutes() chi.Router {
	TokenRouter := chi.NewRouter()
	tokenDb := db.GetTokenDb()

	TokenRouter.Get("/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		struserID := chi.URLParam(r, "user_id")
		userID, err := strconv.Atoi(struserID)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}

		tokens, err := tokenDb.GetUserTokens(int64(userID))
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(tokens)
	})

	TokenRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		TokenReq := new(TokenRequest)
		err := json.NewDecoder(r.Body).Decode(&TokenReq)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}

		UserId, err := strconv.Atoi(TokenReq.UserID)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}

		tokens, err := tokenDb.GetToken(int64(UserId), TokenReq.Provider)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(tokens)
	})

	TokenRouter.Post("/create/", func(w http.ResponseWriter, r *http.Request) {
		CreateReq := new(TokenCreateRequest)
		var NewToken models.Token
		err := json.NewDecoder(r.Body).Decode(&CreateReq)

		if err != nil {
			w.WriteHeader(401)
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
				w.WriteHeader(401)
				w.Write([]byte(err.Error()))
				return
			}
			json.NewEncoder(w).Encode(&token)
		} else {
			Oldtoken.AccessToken = CreateReq.Token
			json.NewEncoder(w).Encode(&Oldtoken)
		}
	})

	return TokenRouter
}
