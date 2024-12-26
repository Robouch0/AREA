//
// EPITECH PROJECT, 2024
// AREA
// File description:
// ioReaderToMap
//

package utils

import (
	"encoding/json"
	"io"
)

func IoReaderToMap(body *io.ReadCloser) (map[string]any, error) {
	b, err := io.ReadAll(*body)
	if err != nil {
		return map[string]any{}, err
	}
	defer (*body).Close()
	var m map[string]any
	if err = json.Unmarshal(b, &m); err != nil {
		return map[string]any{}, err
	}
	return m, nil
}
