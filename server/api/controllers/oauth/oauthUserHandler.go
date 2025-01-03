//
// EPITECH PROJECT, 2024
// AREA
// File description:
// oauthUserHandler
//

package oauth

import (
	"area/api/controllers/log_types"
	"area/api/middleware"
	"area/db"
	"area/models"
	http_utils "area/utils/httpUtils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func CreateUserWithEmail(w http.ResponseWriter, email string, JwtTok *jwtauth.JWTAuth) (*models.User, error) {
	userDb := db.GetUserDb()
	us := new(models.User)

	errUserNotFound := userDb.Db.NewSelect().
		Model(us).
		Where("email = ?", email).
		Scan(context.Background())
	if errUserNotFound != nil {
		var newUser models.User
		newUser.Email = email
		_, err := userDb.CreateUser(&newUser)

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return nil, err
		}
		us = &newUser
	}
	b, err := json.Marshal(log_types.UserLogInfos{Token: middleware.CreateToken(JwtTok, us.ID), UserID: us.ID})
	if err != nil {
		http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
		return nil, err
	}
	w.WriteHeader(200)
	w.Write(b)
	return us, nil
}

func CreateToken(w http.ResponseWriter, user *models.User, AccessToken string, Service string) {
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
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
	} else {
		_, err := tokenDb.UpdateUserTokenByProvider(int64(user.ID), Service, AccessToken)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
