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

func GetUserDb() *UserDb {
	db := initDB()

	return &UserDb{Db: db}
} // Maybe create the table

func (user *UserDb) CreateUser(userData *models.User) (*models.User, error) {
	_, err := user.Db.NewInsert().Model(userData).Exec(context.TODO()) // Do the query

	if err != nil {
		return nil, err
	}
	return userData, nil
}
