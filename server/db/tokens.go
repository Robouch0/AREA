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
		return nil, err
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
