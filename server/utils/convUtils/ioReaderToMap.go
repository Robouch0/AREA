//
// EPITECH PROJECT, 2024
// AREA
// File description:
// ioReaderToMap
//

package conv_utils

import (
	"encoding/json"
	"io"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IoReaderToMap(body *io.ReadCloser) (map[string]any, error) {
	if body == nil {
		return map[string]any{}, status.Errorf(codes.InvalidArgument, "io readCloser is nil")
	}
	if *body == http.NoBody {
		return map[string]any{}, nil
	}
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
