//
// EPITECH PROJECT, 2024
// AREA
// File description:
// main
//

package main

import (
	"log"
	"net/http"

	"area/api/routes"
	"area/gRPC/grpc_routes"
	_ "area/docs"
)

func main() {
	go grpc_routes.LaunchServices()
	apiGateway, err := routes.InitHTTPServer()

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("API Gateway server listening at :3000")
	http.ListenAndServe(":3000", apiGateway.Router)
}
