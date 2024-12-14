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
	"area/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type userSignUp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Email    string `json:"email"`
	Password string `json:"password"`
}

// Account godoc
// @Summary      Sign-up a new account
// @Description  register an account by giving credentials
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param 		 newUser body	models.User	true 	"New User informations to sign-up to the app"
// @Success      200  {object}  userSignUp
// @Failure      401  {object}  error
// @Failure      500  {object}  error
// @Router       /sign-up/ [post]
func SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser userSignUp

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		utils.WriteHTTPResponseErr(&w, 401, "Invalid request body")
		log.Println(err)
		return
	}

	userDb := db.GetUserDb()
	us := new(models.User)
	err = userDb.Db.NewSelect().
		Model(us).
		Where("email = ?", us.Email).
		Scan(context.Background())
	if err == nil {
		utils.WriteHTTPResponseErr(&w, 401, "An user already exists with this email address")
		log.Println(err)
		return
	}

	res, err := userDb.CreateUser(&models.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
	})
	if err != nil {
		utils.WriteHTTPResponseErr(&w, 500, "Error while creating an user")
		log.Println(err)
		return
	}

	b, err := json.Marshal(res)
	if err != nil {
		utils.WriteHTTPResponseErr(&w, 500, "Error while converting results to json")
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(b)
}
