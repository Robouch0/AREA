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

	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	userRouter := chi.NewRouter()
	userDb := db.GetUserDb()

	userRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := userDb.CreateUser(&models.User{ID: 1, FirstName: "Pahul"})

		if err == nil {
			return
		}

		err = json.NewEncoder(w).Encode(res)
		if err == nil {
			return
		}
	})
	return userRouter
}
