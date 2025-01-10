//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package grpc_routes

import (
	asana_server "area/gRPC/api/asana/asanaServer"
	dateTime_server "area/gRPC/api/dateTime/dateTimeServer"
	discord_server "area/gRPC/api/discord/discordServer"
	github "area/gRPC/api/github/githubServer"
	gitlab_server "area/gRPC/api/gitlab/gitlabServer"
	google_server "area/gRPC/api/google/googleServer"
	huggingFace_server "area/gRPC/api/hugging_face/hugging_faceServer"
	miro_server "area/gRPC/api/miro/miroServer"
	reaction_server "area/gRPC/api/reaction/reactionServer"
	spotify_server "area/gRPC/api/spotify/spotifyServer"
	weather_server "area/gRPC/api/weather/weatherServer"
	services "area/protogen/gRPC/proto"
	"cmp"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func LaunchServices() {
	const addr = "localhost:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	dtService, errDt := dateTime_server.NewDateTimeService()
	reactService, errReact := reaction_server.NewReactionService()
	huggingFaceService, errHf := huggingFace_server.NewHuggingFaceService()
	githubService, errGit := github.NewGithubService()
	gitlabService, errGitlab := gitlab_server.NewGitlabService()
	discordService, errDiscord := discord_server.NewDiscordService()
	googleService, errGoogle := google_server.NewGoogleService()
	spotifyService, errSpotify := spotify_server.NewSpotifyService()
	weatherService, errWeather := weather_server.NewWeatherService()
	miroService, errMiro := miro_server.NewMiroService()
	asanaService, errAsana := asana_server.NewAsanaService()

	if err = cmp.Or(errDt, errReact, errGit, errHf, errGoogle, errDiscord, errSpotify, errGitlab, errWeather, errAsana, errMiro); err != nil {
		log.Println(err)
		return
	}
	services.RegisterDateTimeServiceServer(s, dtService)
	services.RegisterHuggingFaceServiceServer(s, huggingFaceService)
	services.RegisterGithubServiceServer(s, githubService)
	services.RegisterGitlabServiceServer(s, gitlabService)
	services.RegisterDiscordServiceServer(s, discordService)
	services.RegisterGoogleServiceServer(s, googleService)
	services.RegisterSpotifyServiceServer(s, spotifyService)
	services.RegisterReactionServiceServer(s, reactService)
	services.RegisterWeatherServiceServer(s, weatherService)
	services.RegisterMiroServiceServer(s, miroService)
	services.RegisterAsanaServiceServer(s, asanaService)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		defer wg.Done()

		log.Printf("gRPC server listening at %v", listener.Addr())
		if err = s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}

	reactService.InitServiceClients(conn)
	dtService.InitReactClient(conn)
	huggingFaceService.InitReactClient(conn)
	googleService.InitReactClient(conn)
	weatherService.InitReactClient(conn)
	githubService.InitReactClient(conn)
	gitlabService.InitReactClient(conn)
	// Init all services with action
	wg.Wait()
}
