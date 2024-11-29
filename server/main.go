//
// EPITECH PROJECT, 2024
// AREA
// File description:
// main
//

package main

import (
	"net/http"

	"area/routes"
)

func main() {
	r := routes.InitHTTPServer()

	http.ListenAndServe(":3000", r)
}
