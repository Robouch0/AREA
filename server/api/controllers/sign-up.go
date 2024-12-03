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
	"net/http"
)

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
	w.WriteHeader(200)
	w.Write(b)
}
