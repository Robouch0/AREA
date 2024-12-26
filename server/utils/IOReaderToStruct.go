//
// EPITECH PROJECT, 2024
// AREA
// File description:
// IOReaderToStruct
//

package utils

import (
	"encoding/json"
	"io"
)

func IOReaderToStruct[T any](body *io.ReadCloser) (*T, error) {
	value := new(T)
	err := json.NewDecoder(*body).Decode(value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
