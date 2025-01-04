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
	grpcutils "area/utils/grpcUtils"
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
// @Description  Get all the tokens of the current logged user
// @Security ApiKeyAuth @Tags         Token
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token [get]
func GetTokens(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := grpcutils.GetUserIDClaim(r.Context())
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
		json.NewEncoder(w).Encode(tokens) // Create a different body
	}
}

// Get token godoc
// @Summary      Get user's token
// @Description  Get the tokens from the current userID and a provider
// @Security ApiKeyAuth
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param 		 provider path	string	true 	"Remote Service Name"
// @Success      200  {object}  models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/{provider} [post]
func GetToken(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		TokenReq := new(TokenInformations)
		TokenReq.Provider = chi.URLParam(r, "provider")

		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}

		tokens, err := tokenDb.GetUserTokenByProvider(int64(userID), TokenReq.Provider)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tokens)
	}
}

// create Token godoc
// @Summary      Create a token
// @Description  Create a token from a user_id and a provider
// @Security ApiKeyAuth
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param 		 tokenCreateRequest body	TokenCreateRequest	true 	"Token creation request informations"
// @Success      200  {object}  models.Token
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/create/ [post]
func CreateTkn(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createReq := new(TokenCreateRequest)
		var newToken models.Token
		err := json.NewDecoder(r.Body).Decode(&createReq)

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}
		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}
		oldToken, err := tokenDb.GetUserTokenByProvider(int64(userID), createReq.Provider)

		if err != nil {
			newToken.AccessToken = createReq.Token
			newToken.Provider = createReq.Provider
			newToken.UserID = int64(userID)
			token, err := tokenDb.CreateToken(&newToken)

			if err != nil {
				http_utils.WriteHTTPResponseErr(&w, 500, err.Error())
				return
			}
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(&token)
		} else {
			oldToken.AccessToken = createReq.Token // That does not work
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(&oldToken)
		}
	}
}

// Delete Token godoc
// @Summary      Delete a token
// @Description  Delete a token from a user_id and a provider
// @Security ApiKeyAuth
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  TokenInformations
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/ [delete]
func DeleteUserToken(tokenDb *db.TokenDb) http.HandlerFunc {
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
