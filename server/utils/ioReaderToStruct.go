//
// EPITECH PROJECT, 2024
// AREA
// File description:
// ioReaderToStruct
//

package utils

import (
	"encoding/json"
	"io"
)

func IoReaderToStruct[T any](body *io.ReadCloser) (*T, error) {
	value := new(T)
	err := json.NewDecoder(*body).Decode(value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
