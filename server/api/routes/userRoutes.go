//
// EPITECH PROJECT, 2024
// AREA
// File description:
// userRoutes
//

package routes

import (
	"area/api/controllers"
	"area/db"

	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	userRouter := chi.NewRouter()
	userDb := db.InitUserDb()

	userRouter.Post("/", controllers.CreateNewUser(userDb))
	// userRouter.Get("/all/", controllers.GetAllUsers(userDb))
	userRouter.Get("/me", controllers.GetUserByID(userDb))

	userRouter.Put("/me", controllers.UpdateUserDatas(userDb))
	return userRouter
}
