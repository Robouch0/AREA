//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-in
//

package controllers

import (
	"area/api/controllers/log_types"
	"area/api/middleware"
	"area/db"
	"area/models"
	"area/utils"
	http_utils "area/utils/httpUtils"
	"context"
	"log"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

// Credentials send by the client for login
type credentials struct {
	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}

// Account godoc
// @Summary      Sign-In
// @Description  Login a user if he has the correct credentials and returns the tokens and the user_id
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param 		 credentials body	credentials	true 	"Credentials of the user who wants to connect"
// @Success      200  {object}  log_types.UserLogInfos
// @Failure      401  {object}  string
// @Failure      500  {object}  string
// @Router       /login/ [post]
func SignIn(jwtauth *jwtauth.JWTAuth) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var cred credentials

		err := json.NewDecoder(r.Body).Decode(&cred)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, "Incorrect body is sent.")
			log.Printf("Json Error: %v\n", err)
			return
		}

		userDb := db.GetUserDb()
		us := new(models.User)
		err = userDb.Db.NewSelect().
			Model(us).
			Where("email = ?", cred.Email).
			Scan(context.Background())

		if !utils.CheckPasswordHash(cred.Password, us.Password) {
			http_utils.WriteHTTPResponseErr(&w, 401, fmt.Sprintf("Invalid password: %s\n", cred.Password))
			log.Printf("Error: %v\n", err)
			return
		}
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, fmt.Sprintf("No user known with email: %s\n", cred.Email))
			log.Printf("Error: %v\n", err)
			return
		}
		b, err := json.Marshal(log_types.UserLogInfos{Token: middleware.CreateToken(jwtauth, us.ID), UserID: us.ID})
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 401, err.Error())
			return
		}
		w.WriteHeader(200)
		w.Write(b)
	}
}
