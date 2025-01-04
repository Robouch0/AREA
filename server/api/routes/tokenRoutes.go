//
// EPITECH PROJECT, 2025
// AREA
// File description:
// tokenRoutes
//

package routes

import (
	"area/api/controllers"
	"area/db"

	"github.com/go-chi/chi/v5"
)

func TokenRoutes() chi.Router {
	tokenRouter := chi.NewRouter()
	tokenDb := db.GetTokenDb()

	tokenRouter.Get("/", controllers.GetTokens(tokenDb))

	tokenRouter.Get("/{provider}", controllers.GetToken(tokenDb))

	tokenRouter.Post("/create/", controllers.CreateTkn(tokenDb))

	tokenRouter.Delete("/{provider}", controllers.DeleteUserToken(tokenDb))

	return tokenRouter
}
