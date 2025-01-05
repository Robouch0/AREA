//
// EPITECH PROJECT, 2024
// AREA
// File description:
// areaRoutes
//

package routes

import (
	"area/api"
	"area/api/controllers/areas"
	"area/db"

	"github.com/go-chi/chi/v5"
)

func AreaRoutes(gateway *api.ApiGateway) chi.Router {
	areaRouter := chi.NewRouter()
	areaDb := db.GetAreaDb()

	areaRouter.Get("/list", areas.GetUserAreas(gateway, areaDb)) // I don't know if more details are necessary
	areaRouter.Put("/activate", areas.ActivateArea(gateway))
	areaRouter.Post("/create/{service}", areas.CreateArea(gateway))
	areaRouter.Get("/create/list", areas.ListService(gateway))
	return areaRouter
}
