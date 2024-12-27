//
// EPITECH PROJECT, 2024
// AREA
// File description:
// users
//

package controllers

import (
	"area/db"
	"area/models"
	"area/utils"
	http_utils "area/utils/httpUtils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserInformations struct {
	ID        uint   `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`
	FirstName string `bun:"first_name" json:"first_name"`
	LastName  string `bun:"last_name" json:"last_name"`

	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}

func UserModelToUserInfos(data *models.User) *UserInformations {
	return &UserInformations{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
	}
}

// Swagooo
func CreateNewUser(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newUser, err := utils.IoReaderToStruct[models.User](&r.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		res, err := userDb.CreateUser(newUser)

		if err != nil {
			fmt.Printf("User Creation Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(UserModelToUserInfos(res)) // We don't send all user info xDD
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
}

func GetAllUsers(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := userDb.GetUsers()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(res) // Convert to []UserInformations
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	}
}

func GetUserByID(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("error converting id to int: %s: %v", strId, err))
			log.Printf("Error: %v\n", err)
			return
		}

		res, err := userDb.GetUser(id)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("error: no such user with id: %d", id))
			log.Printf("Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(UserModelToUserInfos(res))
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("Error: %v\n", err))
			log.Printf("Error: %v\n", err)
			return
		}
	}
}

func UpdateUserDatas(userDb *db.UserDb) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("error converting id to int: %s: %v", strId, err))
			return
		}
		updateData, err := utils.IoReaderToStruct[models.UpdatableUserData](&r.Body)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("Invalid request format"))
			return
		}

		_, err = userDb.UpdateUserData(id, updateData)
		if err != nil {
			http_utils.WriteHTTPResponseErr(&w, 422, fmt.Sprintf("Invalid informations sent to update user"))
			log.Println(err)
			return
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(updateData)
	}
}
