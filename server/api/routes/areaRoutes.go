//
// EPITECH PROJECT, 2024
// AREA
// File description:
// areaRoutes
//

package routes

import (
	"area/api"
	"area/api/controllers"
	"area/db"

	"github.com/go-chi/chi/v5"
)

func AreaRoutes(gateway *api.ApiGateway) chi.Router {
	areaRouter := chi.NewRouter()
	areaDb := db.GetAreaDb()

	areaRouter.Get("/list", controllers.GetUserAreas(gateway, areaDb))
	return areaRouter
}
