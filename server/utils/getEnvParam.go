//
// EPITECH PROJECT, 2024
// AREA
// File description:
// getEnvParam
//

package utils

import (
	"errors"

	"github.com/spf13/viper"
)

func GetEnvParameter(key string) (string, error) {
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}
	tok := viper.Get(key)
	if tok == nil {
		return "", errors.New("no such environment variable")
	}
	return tok.(string), nil
}

func GetEnvParameterToBearer(key string) (string, error) {
	tok, err := GetEnvParameter(key)

	if err != nil {
		return "", err
	}
	return "Bearer " + tok, nil
}
