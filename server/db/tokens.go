//
// EPITECH PROJECT, 2024
// AREA
// File description:
// OAuth
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenDb struct {
	Db *bun.DB
}

func InitTokenDb() (*TokenDb, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Token)(nil)).
		IfNotExists().
		Exec(context.Background())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot init token database %v", err)
	}

	return &TokenDb{Db: db}, nil
}

func GetTokenDb() *TokenDb {
	db := initDB()
	return &TokenDb{Db: db}
}

func (token *TokenDb) CreateToken(newToken *models.Token) (*models.Token, error) {
	_, err := token.Db.NewInsert().
		Model(newToken).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func (token *TokenDb) UpdateUserTokenByProvider(userID int64, provider string, accessToken string) (*models.Token, error) {
	upTok := new(models.Token)
	_, err := token.Db.NewUpdate().
		Model(upTok).
		Set("access_token = ?", accessToken).
		Where("user_id = ?", userID).
		Where("provider = ?", provider).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return upTok, nil
}

func (Token *TokenDb) GetTokens() (*([]models.Token), error) {
	allTokens := new([]models.Token)
	err := Token.Db.NewSelect().
		Model(allTokens).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allTokens, nil
}

func (Token *TokenDb) GetUserTokens(userID int64) (*([]models.Token), error) {
	allTokens := new([]models.Token)
	err := Token.Db.NewSelect().
		Model(allTokens).
		Where("user_id = ?", userID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allTokens, nil
}

func (Token *TokenDb) GetUserTokenByProvider(userID int64, provider string) (*models.Token, error) {
	us := new(models.Token)

	err := Token.Db.NewSelect().
		Model(us).
		Where("user_id = ? AND provider = ?", userID, provider).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (Token *TokenDb) DeleteUserTokenByProvider(userID int64, provider string) (*models.Token, error) {
	us := new(models.Token)
	_, err := Token.Db.NewDelete().
		Model(us).
		Where("user_id = ? AND provider = ?", userID, provider).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return us, nil
}
