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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func InitHTTPServer() (*api.ApiGateway, error) {
	gateway, err := api.CreateApiGateway()
	if err != nil {
		return nil, err
	}

	gateway.Router.Use(middleware.Logger)
	gateway.Router.Use(middleware.AllowContentType("application/json"))
	gateway.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		MaxAge:         300,
	}))

	gateway.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	gateway.Router.Get("/about.json", controllers.AboutRoute(gateway))

	gateway.Router.Mount("/oauth/", controllers.OAuthRoutes(gateway.JwtTok))
	gateway.Router.Post("/login/", controllers.SignIn(gateway.JwtTok))
	gateway.Router.Post("/sign-up/", controllers.SignUp)

	gateway.Router.Mount("/webhook/", controllers.WebHookRoutes(gateway))

	gateway.Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(gateway.JwtTok))
		r.Use(jwtauth.Authenticator(gateway.JwtTok))

		r.Mount("/token/", TokenRoutes())
		r.Mount("/user/", UserRoutes())
		r.Mount("/areas/", AreaRoutes(gateway))
		r.Get("/ping", controllers.PingRoute)
	})
	return gateway, nil
}
