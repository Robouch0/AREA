//
// EPITECH PROJECT, 2024
// AREA
// File description:
// decodeBase64ToJSON
//

package utils

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DecodeBase64ToStruct[T any](data []byte) (*T, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	log.Println(string(decoded))
	structData := new(T)
	if json.Unmarshal(decoded, structData) != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Could not decode data sent")
	}
	return structData, nil
}
