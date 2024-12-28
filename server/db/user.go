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

func InitUserDb() *UserDb {
	db := initDB()

	db.NewCreateTable().
		Model((*models.User)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &UserDb{Db: db}
}

func GetUserDb() *UserDb {
	db := initDB()
	return &UserDb{Db: db}
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

func (user *UserDb) GetUsers() (*([]models.User), error) { // Use the generic function
	allUsers := new([]models.User)
	err := user.Db.NewSelect().
		Model(allUsers).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allUsers, nil
}

func (user *UserDb) GetUserByID(id int) (*models.User, error) { // Use the generic function
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

func (user *UserDb) UpdateUserData(id int, datas *models.UpdatableUserData) (*models.User, error) {
	userData := new(models.User)
	_, err := user.Db.NewUpdate().
		Model(userData).
		Set("first_name = ?", datas.FirstName).
		Set("last_name = ?", datas.LastName).
		Set("password = ?", datas.Password).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return userData, nil
}
