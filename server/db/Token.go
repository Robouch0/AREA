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

type OAuthDb struct {
	Db *bun.DB
}

func InitOAuthDb() *OAuthDb {
	db := initDB()

	db.NewCreateTable().
		Model((*models.Token)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &OAuthDb{Db : db}
}

func GetOAuthDb() *OAuthDb {
	db := initDB()
	return &OAuthDb{Db: db}
}

func (OAuth *OAuthDb) CreateOAuthToken(token *models.Token) (*models.Token, error) {
	_, err := OAuth.Db.NewInsert().
		Model(token).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (OAuth *OAuthDb) GetOAuthTokens() (*([]models.Token), error) {
	allTokens := new([]models.Token)
	err := OAuth.Db.NewSelect().
		Model(allTokens).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allTokens, nil
}

func (OAuth *OAuthDb) getOAuthToken(userID int64, provider string) (*models.Token, error) {
	us := new(models.Token)

	err := OAuth.Db.NewSelect().
	Model(us).
	Where("user_id = ? AND provider = ?", userID, provider).
	Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return us, nil
}
