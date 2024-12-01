//
// EPITECH PROJECT, 2024
// AREA
// File description:
// sign-in
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

type Credentials struct {
	UserName string `bun:"user_name" json:"user_name"`
	Password string `bun:"password" json:"password"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(401)
		fmt.Printf("Error: %v\n", err)
		return
	}
	userDb := db.GetUserDb()

	us := new(models.User)
	err = userDb.Db.NewSelect().
		Model(us).
		Where("first_name = ?", cred.UserName). // The check has to be better
		Scan(context.Background())

	if err != nil {
		w.WriteHeader(401)
		fmt.Printf("Error: %v\n", err)
		return
	}
	w.Write([]byte("Login done"))
}
