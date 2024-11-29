//
// EPITECH PROJECT, 2024
// AREA
// File description:
// userRoutes
//

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct { // will be a postresql model maybe
	ID   string `json:"id"`
	NAME string `json:"name"`
}

func UserRoutes() chi.Router {
	userRouter := chi.NewRouter()

	userRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(User{ID: "1", NAME: "Pahul"})
		if err == nil {
			return
		}
	})
	return userRouter
}
