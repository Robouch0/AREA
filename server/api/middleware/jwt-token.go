//
// EPITECH PROJECT, 2024
// AREA
// File description:
// jwt-token
//

package middleware

import (
	"area/utils"

	"github.com/go-chi/jwtauth/v5"
)

func GetNewJWTAuth() *jwtauth.JWTAuth {
	tok, err := utils.GetEnvParameterToBearer("SECRET_KEY")
	if err != nil {
		return nil
	}
	return jwtauth.New("HS256", []byte(tok), nil)
}

func CreateToken(jwtauth *jwtauth.JWTAuth, userId uint) string {
	_, tokenString, _ := jwtauth.Encode(map[string]interface{}{"user_id": userId})
	return tokenString
}
