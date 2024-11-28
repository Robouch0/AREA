//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type User struct {
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

func InitHTTPServer() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Mount("/users/", UserRoutes())
	return r
}
