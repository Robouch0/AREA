//
// EPITECH PROJECT, 2024
// AREA
// File description:
// userRoutes
//

package routes

import (
	"area/db"
	"area/models"
	"fmt"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	userRouter := chi.NewRouter()
	userDb := db.InitUserDb()

	userRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		newUser := new(models.User)
		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		res, err := userDb.CreateUser(newUser)

		if err != nil {
			fmt.Printf("User Creation Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	})

	userRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := userDb.GetUsers()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	})

	userRouter.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(fmt.Sprintf("error converting id to int: %s: %v", strId, err)))
			return
		}

		res, err := userDb.GetUserByID(id)
		if err != nil {
			w.WriteHeader(422)
			w.Write([]byte(fmt.Sprintf("error: no such user with id: %d", id)))
			return
		}

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	})
	return userRouter
}
