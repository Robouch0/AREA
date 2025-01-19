//
// EPITECH PROJECT, 2024
// AREA
// File description:
// ioReaderToStruct
//

package conv_utils

import (
	"encoding/json"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IoReaderToStruct[T any](body *io.ReadCloser) (*T, error) {
	if body == nil {
		return nil, status.Errorf(codes.InvalidArgument, "io readCloser is nil")
	}
	value := new(T)
	err := json.NewDecoder(*body).Decode(value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
