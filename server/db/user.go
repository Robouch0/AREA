//
// EPITECH PROJECT, 2024
// AREA
// File description:
// user
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type UserDb struct {
	Db *bun.DB
}

type OAuthDb struct {
	Db *bun.DB
}

func InitUserDb() *UserDb {
	db := initDB()

	db.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &UserDb{Db: db}
}

func InitOAuthDb() *OAuthDb {
	db := initDB()

	db.NewCreateTable().
		Model((*models.OAuthToken)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &OAuthDb{Db : db}
}

func GetUserDb() *UserDb {
	db := initDB()
	return &UserDb{Db: db}
}

func GetOAuthDb() *OAuthDb {
	db := initDB()
	return &OAuthDb{Db: db}
}

func (OAuth *OAuthDb) CreateOAuthToken(token *models.OAuthToken) (*models.OAuthToken, error) {
	_, err := OAuth.Db.NewInsert().
		Model(token).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (OAuth *OAuthDb) GetOAuthTokens() (*([]models.OAuthToken), error) {
	allTokens := new([]models.OAuthToken)
	err := OAuth.Db.NewSelect().
		Model(allTokens).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allTokens, nil
}

func (OAuth *OAuthDb) getOAuthToken(userID int64, provider string) (*models.OAuthToken, error) {
	us := new(models.OAuthToken)

	err := OAuth.Db.NewSelect().
	Model(us).
	Where("user_id = ? AND provider = ?", userID, provider).
	Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (user *UserDb) CreateUser(userData *models.User) (*models.User, error) {
	_, err := user.Db.NewInsert().
		Model(userData).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (user *UserDb) GetUsers() (*([]models.User), error) {
	allUsers := new([]models.User)
	err := user.Db.NewSelect().
		Model(allUsers).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allUsers, nil
}

func (user *UserDb) GetUser(id int) (*models.User, error) {
	us := new(models.User)
	err := user.Db.NewSelect().
		Model(us).
		Where("id = ?", id).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return us, nil
}
