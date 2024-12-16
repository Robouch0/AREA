//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-up
//

package controllers

import (
	"area/db"
	"area/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Account godoc
// @Summary      Sign-up a new account
// @Description  register an account by giving credentials
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param 		 newUser body	models.User	true 	"New User informations to sign-up to the app"
// @Success      200  {object}  models.User
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /sign-up/ [post]
func SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(401)
		fmt.Printf("Json Error: %v\n", err)
		return
	}

	userDb := db.GetUserDb()
	us := new(models.User)
	err = userDb.Db.NewSelect().
		Model(us).
		Where("email = ?", us.Email).
		Scan(context.Background())
	if err == nil {
		w.WriteHeader(401)
		w.Write([]byte("An user already exists with this email address"))
		return
	}

	res, err := userDb.CreateUser(&newUser)
	if err != nil {
		w.WriteHeader(500)
		fmt.Println(err)
		w.Write([]byte("Error while creating an user"))
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error while converting results to json"))
		return
	}

	log.Println("New user: ", res)
	w.WriteHeader(200)
	w.Write(b)
}
