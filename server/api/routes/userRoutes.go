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
	userRouter.Get("/", controllers.GetAllUsers(userDb))
	userRouter.Get("/{id}", controllers.GetUserByID(userDb))

	userRouter.Put("/{id}", controllers.UpdateUserDatas(userDb))
	return userRouter
}
