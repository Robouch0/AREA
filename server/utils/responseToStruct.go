//
// EPITECH PROJECT, 2024
// AREA
// File description:
// responseToStruct
//

package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseToStruct[T any](resp *http.Response) (*T, error) {
	value := new(T)
	err := json.NewDecoder(resp.Body).Decode(value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
