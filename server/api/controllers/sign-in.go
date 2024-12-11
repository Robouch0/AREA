//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-in
//

package controllers

import (
	"area/api/middleware"
	"area/db"
	"area/models"
	"context"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

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
// @Success      200  {object}  string
// @Failure      401  {object}  string
// @Failure      500  {object}  string
// @Router       /login/ [post]
func SignIn(jwtauth *jwtauth.JWTAuth) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var cred credentials

		err := json.NewDecoder(r.Body).Decode(&cred)
		if err != nil {
			w.WriteHeader(401)
			fmt.Printf("Json Error: %v\n", err)
			return
		}
		userDb := db.GetUserDb()

		us := new(models.User)
		err = userDb.Db.NewSelect().
			Model(us).
			Where("email = ?", cred.Email).
			Where("password = ?", cred.Password).
			Scan(context.Background())

		if err != nil {
			w.WriteHeader(401)
			fmt.Printf("Error: %v\n", err)
			return
		}
		w.Write([]byte(fmt.Sprintf("%s,%d", middleware.CreateToken(jwtauth, us.ID), us.ID)))
	}
}
