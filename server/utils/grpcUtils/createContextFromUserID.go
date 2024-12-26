//
// EPITECH PROJECT, 2024
// AREA
// File description:
// createContextFromUserID
//

package grpcutils

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

func CreateContextFromUserID(userID int) context.Context {
	md := metadata.Pairs("user_id", strconv.Itoa(userID))
	return metadata.NewOutgoingContext(context.Background(), md)
}
