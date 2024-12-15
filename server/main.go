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
	_ "area/docs"
	"area/gRPC/grpc_routes"
)

func main() {
	go grpc_routes.LaunchServices()
	apiGateway, err := routes.InitHTTPServer()

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("API Gateway server listening at :8080")
	http.ListenAndServe(":8080", apiGateway.Router)
}
