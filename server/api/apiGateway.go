//
// EPITECH PROJECT, 2024
// AREA
// File description:
// apiGateway
//

package api

import (
	areaMiddleware "area/api/middleware"
	"area/gRPC/api/dateTime"
	"area/gRPC/api/discord"
	"area/gRPC/api/github"
	gitlab_client "area/gRPC/api/gitlab/gitlabClient"
	google_client "area/gRPC/api/google/googleClient"
	huggingFace "area/gRPC/api/hugging_face"
	"area/gRPC/api/reaction"
	IServ "area/gRPC/api/serviceInterface"
	"area/gRPC/api/spotify"
	weather_client "area/gRPC/api/weather/weatherClient"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiGateway struct {
	Router chi.Router
	JwtTok *jwtauth.JWTAuth

	conn    *grpc.ClientConn
	Clients map[string]IServ.ClientService
}

func CreateApiGateway() (*ApiGateway, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	m := make(map[string]IServ.ClientService)
	m["dt"] = dateTime.NewDateTimeServiceClient(conn)
	m["hf"] = huggingFace.NewHuggingFaceClient(conn)
	m["github"] = github.NewGithubClient(conn)
	m["gitlab"] = gitlab_client.NewGitlabClient(conn)
	m["google"] = google_client.NewGoogleClient(conn)
	m["discord"] = discord.NewDiscordClient(conn)
	m["react"] = reaction.NewReactionServiceClient(conn)
	m["spotify"] = spotify.NewSpotifyClient(conn)
	m["weather"] = weather_client.NewWeatherClient(conn)
	return &ApiGateway{
		Router:  chi.NewRouter(),
		JwtTok:  areaMiddleware.GetNewJWTAuth(),
		conn:    conn,
		Clients: m,
	}, nil
}
