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

func convertTokensToPublicInfos(tokens *[]models.Token) []TokenInformations {
	var allTokens []TokenInformations

	for _, tok := range *tokens {
		allTokens = append(allTokens, TokenInformations{
			UserID:   strconv.FormatInt(tok.UserID, 10),
			Provider: tok.Provider,
		})
	}
	return allTokens
}

// Get tokens godoc
// @Summary      Get all the tokens from a user
// @Description  Get all the tokens of the current logged user
// @Security ApiKeyAuth
// @Tags         Token
// @Accept       json
// @Produce      json
// @Success      200  {object}  []TokenInformations
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
		allToks := convertTokensToPublicInfos(tokens)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&allToks)
	}
}

// Get token godoc
// @Summary      Get user's token
// @Description  Get a the token associated to the remote provider of the user
// @Security ApiKeyAuth
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param 		 provider path	string	true 	"Remote Service Name"
// @Success      200  {object}  TokenInformations
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/{provider} [get]
func GetToken(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenReq := new(TokenInformations)
		tokenReq.Provider = chi.URLParam(r, "provider")

		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}

		token, err := tokenDb.GetUserTokenByProvider(int64(userID), tokenReq.Provider)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&TokenInformations{
			UserID:   strconv.FormatInt(token.UserID, 10),
			Provider: token.Provider,
		})
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
// @Success      200  {object}  TokenInformations
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
			json.NewEncoder(w).Encode(&TokenInformations{
				UserID:   strconv.FormatInt(int64(userID), 10),
				Provider: token.Provider,
			})
		} else {
			tokenDb.UpdateUserTokenByProvider(int64(userID), createReq.Provider, createReq.Token)
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(&TokenInformations{
				UserID:   strconv.FormatInt(int64(userID), 10),
				Provider: oldToken.Provider,
			})
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
// @Param 		 provider path	string	true 	"Remote Service Name"
// @Success      200  {object}  TokenInformations
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /token/ [delete]
func DeleteUserToken(tokenDb *db.TokenDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")

		userID, err := grpcutils.GetUserIDClaim(r.Context())
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 400, err.Error())
			return
		}

		_, err = tokenDb.DeleteUserTokenByProvider(int64(userID), provider)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Error while deleting token")
			log.Println(err)
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(&TokenInformations{
			UserID:   strconv.FormatInt(int64(userID), 10),
			Provider: provider,
		})
	}
}
