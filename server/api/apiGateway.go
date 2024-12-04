//
// EPITECH PROJECT, 2024
// AREA
// File description:
// apiGateway
//

package api

import (
	areaMiddleware "area/api/middleware"
	"area/gRPC/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiGateway struct {
	Router chi.Router
	JwtTok *jwtauth.JWTAuth

	conn    *grpc.ClientConn
	Clients map[string]api.ClientService
}

func CreateApiGateway() (*ApiGateway, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	m := make(map[string]api.ClientService)
	m["hello"] = api.NewHelloServiceClient(conn)
	m["dt"] = api.NewDateTimeServiceClient(conn)
	m["react"] = api.NewReactionServiceClient(conn)
	return &ApiGateway{
		Router:  chi.NewRouter(),
		JwtTok:  areaMiddleware.GetNewJWTAuth(),
		conn:    conn,
		Clients: m,
	}, nil
}
