//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package routes

import (
	api "area/api"
	"area/api/controllers"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"

	_ "area/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger AREA API
// @version 1.0
// @description This is a the document of the Backend routes of the application AREA
func InitHTTPServer() (*api.ApiGateway, error) {
	gateway, err := api.CreateApiGateway()
	if err != nil {
		return nil, err
	}

	gateway.Router.Use(middleware.Logger)
	gateway.Router.Use(middleware.AllowContentType("application/json"))

	gateway.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	gateway.Router.Get("/ping", controllers.PingRoute)
	gateway.Router.Get("/about.json", controllers.AboutRoute)

	gateway.Router.Mount("/users/", UserRoutes())
	gateway.Router.Mount("/oauth/", controllers.OAuthRoutes(gateway.JwtTok))
	gateway.Router.Mount("/token/", controllers.TokenRoutes())

	gateway.Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(gateway.JwtTok))
		r.Use(jwtauth.Authenticator(gateway.JwtTok))

		r.Get("/admin/", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context()) // DO NOT FORGET TO CHECK THE CLAIM
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})

		r.Post("/create/{service}", controllers.CreateRoute(gateway))
		r.Get("/create/list", controllers.ListService(gateway))
		r.Get("/ping", controllers.PingRoute)
	})
	gateway.Router.Post("/login/", controllers.SignIn(gateway.JwtTok))
	gateway.Router.Post("/sign-up/", controllers.SignUp)
	return gateway, nil
}
