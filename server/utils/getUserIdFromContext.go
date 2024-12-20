//
// EPITECH PROJECT, 2024
// AREA
// File description:
// getUserIdFromContext
//

package utils

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-chi/jwtauth/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetUserIdFromContext(ctx context.Context, serviceName string) (uint, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Errorf(codes.DataLoss, fmt.Sprintf("%v: Failed to get metadata", serviceName))
	}
	if userIDString, ok := md["user_id"]; ok {
		if len(userIDString) == 0 {
			return 0, status.Errorf(codes.DataLoss, fmt.Sprintf("%v: Failed to get UserID claim data", serviceName))
		}
		if userID, errAtoi := strconv.Atoi(userIDString[0]); errAtoi != nil {
			return 0, nil
		} else {
			return uint(userID), nil
		}
	}
	return 0, status.Errorf(codes.DataLoss, fmt.Sprintf("%v: No user Id in claim data", serviceName))
}

func GetUserIDClaim(ctx context.Context) (uint, error) {
	_, claims, _ := jwtauth.FromContext(ctx)

	userIDClaims, ok := (claims["user_id"].(float64))
	if !ok {
		return 0, errors.New("Cannot get the user ID from the JWT claims")
	}
	return uint(userIDClaims), nil
}
