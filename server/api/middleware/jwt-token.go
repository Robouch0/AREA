//
// EPITECH PROJECT, 2024
// AREA
// File description:
// jwt-token
//

package middleware

import "github.com/go-chi/jwtauth/v5"

const SECRET_KEY = "<jwt-secret-env>"

func GetNewJWTAuth() *jwtauth.JWTAuth {
	return jwtauth.New("HS256", []byte(SECRET_KEY), nil)
}

func CreateToken(jwtauth *jwtauth.JWTAuth, userId uint) string {
	_, tokenString, _ := jwtauth.Encode(map[string]interface{}{"user_id": userId})
	return tokenString
}
