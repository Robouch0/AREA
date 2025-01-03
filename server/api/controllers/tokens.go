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
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TokenInformations struct {
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
// @Param 		 user_id path	string	true 	"Id of the user"
// @Success      200  {object}  []models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/{user_id} [get]
func getTokens(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		struserID := chi.URLParam(r, "user_id")
		userID, err := strconv.Atoi(struserID)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}

		tokens, err := tokenDb.GetUserTokens(int64(userID))
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 500, err.Error())
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
// @Param 		 user_id path	string	true 	"Id of the user"
// @Param 		 provider path	string	true 	"Provider of the Remote Service"
// @Success      200  {object}  models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/ [post]
func getToken(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		TokenReq := new(TokenInformations)
		TokenReq.UserID = chi.URLParam(r, "user_id")
		TokenReq.Provider = chi.URLParam(r, "provider")

		UserId, err := strconv.Atoi(TokenReq.UserID)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		tokens, err := tokenDb.GetUserTokenByProvider(int64(UserId), TokenReq.Provider)
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
		Oldtoken, err := tokenDb.GetUserTokenByProvider(int64(UserId), CreateReq.Provider)

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

// Delete Token godoc
// @Summary      Delete a token
// @Description  Delete a token from a user_id and a provider
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  TokenInformations
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/ [delete]
func deleteUserToken(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userStrID := chi.URLParam(r, "user_id")
		provider := chi.URLParam(r, "provider")

		userId, err := strconv.Atoi(userStrID)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Invalid userID given")
			log.Println(err)
			return
		}

		_, err = tokenDb.DeleteUserTokenByProvider(int64(userId), provider)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Error while deleting token")
			log.Println(err)
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&TokenInformations{UserID: userStrID, Provider: provider})
	}
}

func TokenRoutes() chi.Router {
	tokenRouter := chi.NewRouter()
	tokenDb := db.GetTokenDb()

	tokenRouter.Get("/{user_id}", getTokens(tokenDb))

	tokenRouter.Get("/{user_id}/{provider}", getToken(tokenDb))

	tokenRouter.Post("/create/", createTkn(tokenDb))

	tokenRouter.Delete("/{user_id}/{provider}", deleteUserToken(tokenDb))

	return tokenRouter
}
