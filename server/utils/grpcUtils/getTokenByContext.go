//
// EPITECH PROJECT, 2024
// AREA
// File description:
// getTokenByContext
//

package grpcutils

import (
	"area/db"
	"area/models"
	"context"
)

func GetTokenByContext(
	ctx context.Context,
	tokenDb *db.TokenDb,
	serviceName string,
	remoteProvider string,
) (*models.Token, error) {
	userID, errClaim := GetUserIdFromContext(ctx, serviceName)
	if errClaim != nil {
		return nil, errClaim
	}

	tokenInfo, err := tokenDb.GetUserTokenByProvider(int64(userID), remoteProvider)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
