//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package routes

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func InitHTTPServer() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	r.Mount("/users/", UserRoutes())
	return r
}
