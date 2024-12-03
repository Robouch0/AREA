//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package routes

import (
	"area/api/controllers"
	areaMiddleware "area/api/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
    "github.com/go-chi/cors"
)

type ApiGateway struct {
	Router chi.Router
	JwtTok *jwtauth.JWTAuth
}

type ClientInfo struct {
	Host string
	Time int64
}

func createApiGateway() *ApiGateway {
	return &ApiGateway{
		Router: chi.NewRouter(),
		JwtTok: areaMiddleware.GetNewJWTAuth(),
	}
}

func InitHTTPServer() *ApiGateway {
	gateway := createApiGateway()
	gateway.Router.Use(middleware.Logger)
	gateway.Router.Use(middleware.AllowContentType("application/json"))

    gateway.Router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
      }))

	gateway.Router.Get("/ping", PingRoute)
	gateway.Router.Get("/about.json", AboutRoute)
	gateway.Router.Mount("/users/", UserRoutes())

	gateway.Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(gateway.JwtTok))
		r.Use(jwtauth.Authenticator(gateway.JwtTok))

		r.Get("/admin/", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context()) // DO NOT FORGET TO CHECK THE CLAIM
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})
	gateway.Router.Post("/login/", controllers.SignIn(gateway.JwtTok))
	gateway.Router.Post("/sign-up/", controllers.SignUp)
	return gateway
}

func PingRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Pong")))
}

func AboutRoute(w http.ResponseWriter, r *http.Request) {
	Clientdata := ClientInfo{
		Host:		r.RemoteAddr,
		Time:	time.Now().Unix(),
	}
	fmt.Printf(r.Header.Get("X-Real-Ip"))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Clientdata)
}
