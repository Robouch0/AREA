//
// EPITECH PROJECT, 2024
// AREA
// File description:
// router
//

package grpc_routes

import (
	"area/gRPC/api/dateTime"
	"area/gRPC/api/discord"
	"area/gRPC/api/github"
	google_server "area/gRPC/api/google/googleServer"
	huggingFace "area/gRPC/api/hugging_face"
	"area/gRPC/api/reaction"
	"area/gRPC/api/spotify"
	services "area/protogen/gRPC/proto"
	"cmp"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func LaunchServices() {
	const addr = "0.0.0.0:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	dtService, errDt := dateTime.NewDateTimeService()
	reactService, errReact := reaction.NewReactionService()
	huggingFaceService, errHf := huggingFace.NewHuggingFaceService()
	githubService, errGit := github.NewGithubService()
	discordService, errDiscord := discord.NewDiscordService()
	googleService, errGoogle := google_server.NewGoogleService()
	spotifyService, errSpotify := spotify.NewSpotifyService()

	if err = cmp.Or(errDt, errReact, errGit, errHf, errGoogle, errDiscord, errSpotify); err != nil {
		log.Println(err)
		return
	}
	services.RegisterDateTimeServiceServer(s, dtService)
	services.RegisterHuggingFaceServiceServer(s, huggingFaceService)
	services.RegisterGithubServiceServer(s, githubService)
	services.RegisterDiscordServiceServer(s, discordService)
	services.RegisterGoogleServiceServer(s, googleService)
	services.RegisterSpotifyServiceServer(s, spotifyService)
	services.RegisterReactionServiceServer(s, reactService)

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
	// Init all services with action
	wg.Wait()
}
