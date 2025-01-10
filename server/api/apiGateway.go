//
// EPITECH PROJECT, 2024
// AREA
// File description:
// apiGateway
//

package api

import (
	areaMiddleware "area/api/middleware"
	asana_client "area/gRPC/api/asana/asanaClient"
	dateTime_client "area/gRPC/api/dateTime/dateTimeClient"
	discord_client "area/gRPC/api/discord/discordClient"
	github "area/gRPC/api/github/githubClient"
	gitlab_client "area/gRPC/api/gitlab/gitlabClient"
	google_client "area/gRPC/api/google/googleClient"
	huggingFace_client "area/gRPC/api/hugging_face/hugging_faceClient"
	miro_client "area/gRPC/api/miro/miroClient"
	reaction_client "area/gRPC/api/reaction/reactionClient"
	IServ "area/gRPC/api/serviceInterface"
	spotify_client "area/gRPC/api/spotify/spotifyClient"
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
	m["dt"] = dateTime_client.NewDateTimeServiceClient(conn)
	m["hf"] = huggingFace_client.NewHuggingFaceClient(conn)
	m["github"] = github.NewGithubClient(conn)
	m["gitlab"] = gitlab_client.NewGitlabClient(conn)
	m["google"] = google_client.NewGoogleClient(conn)
	m["discord"] = discord_client.NewDiscordClient(conn)
	m["react"] = reaction_client.NewReactionServiceClient(conn)
	m["spotify"] = spotify_client.NewSpotifyClient(conn)
	m["weather"] = weather_client.NewWeatherClient(conn)
	m["miro"] = miro_client.NewMiroClient(conn)
	m["asana"] = asana_client.NewAsanaClient(conn)
	return &ApiGateway{
		Router:  chi.NewRouter(),
		JwtTok:  areaMiddleware.GetNewJWTAuth(),
		conn:    conn,
		Clients: m,
	}, nil
}
