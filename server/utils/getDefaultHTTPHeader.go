//
// EPITECH PROJECT, 2024
// AREA
// File description:
// getDefaultHTTPHeader
//

package utils

import "net/http"

func GetDefaultHTTPHeader(bearerTok string) http.Header {
	header := http.Header{}

	header.Set("Authorization", bearerTok)
	header.Add("Content-Type", "application/json;charset=UTF-8")
	return header
}

func GetDefaultBearerHTTPHeader(token string) http.Header {
	return GetDefaultHTTPHeader("Bearer " + token)
}
