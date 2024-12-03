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
	"area/gRPC/api"
	"fmt"
	"log"

	// "fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiGateway struct {
	Router chi.Router
	JwtTok *jwtauth.JWTAuth

	conn    *grpc.ClientConn
	clients map[string]api.ClientService
}

func createApiGateway() (*ApiGateway, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	m := make(map[string]api.ClientService)
	m["hello"] = api.NewHelloServiceClient(conn)
	return &ApiGateway{
		Router:  chi.NewRouter(),
		JwtTok:  areaMiddleware.GetNewJWTAuth(),
		conn:    conn,
		clients: m,
	}, nil
}

func InitHTTPServer() (*ApiGateway, error) {
	gateway, err := createApiGateway()
	if err != nil {
		return nil, err
	}

	gateway.Router.Use(middleware.Logger)
	gateway.Router.Use(middleware.AllowContentType("application/json"))

	gateway.Router.Mount("/users/", UserRoutes())

	gateway.Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(gateway.JwtTok))
		r.Use(jwtauth.Authenticator(gateway.JwtTok))

		r.Get("/admin/", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context()) // DO NOT FORGET TO CHECK THE CLAIM
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})

		r.Post("/create/{service}", func(w http.ResponseWriter, r *http.Request) {
			serviceParam := chi.URLParam(r, "service")

			if service, ok := gateway.clients[serviceParam]; ok {
				msg, err := service.SendAction(serviceParam)
				if err != nil {
					log.Println(err)
					return
				}
				w.Write([]byte(msg))

			} else {
				w.WriteHeader(401)
				w.Write([]byte(fmt.Sprintf("No such Service: %v", serviceParam)))
			}
		})
	})
	gateway.Router.Post("/login/", controllers.SignIn(gateway.JwtTok))
	gateway.Router.Post("/sign-up/", controllers.SignUp)
	return gateway, nil
}
