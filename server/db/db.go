//
// EPITECH PROJECT, 2024
// AREA
// File description:
// db
//

package db

import (
	"area/utils"
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	dbConnOnce sync.Once
)

func GetAll[T any](Db bun.IDB) (*[]T, error) {
	allDatas := new([]T)
	err := Db.NewSelect().
		Model(allDatas).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func GetByID[T any](Db bun.IDB, ID uint) (*T, error) {
	allDatas := new(T)
	err := Db.NewSelect().
		Model(allDatas).
		Where("id = ?", ID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func GetByActionID[T any](Db bun.IDB, action_id uint) (*T, error) {
	allDatas := new(T)
	err := Db.NewSelect().
		Model(allDatas).
		Where("action_id = ?", action_id).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func SetActivateByActionID[T any](Db bun.IDB, activated bool, userID, actionID uint) (*T, error) {
	data := new(T)
	_, err := Db.NewUpdate().
		Model(data).
		Set("activated = ?", activated).
		Where("user_id = ?", userID).
		Where("action_id = ?", actionID).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	data, err = GetByActionID[T](Db, actionID)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Could not set activated for time action")
	}
	return data, nil
}

func DeleteUserActionByActionID[T any](Db bun.IDB, userID, actionID uint) error {
	data := new(T)
	_, err := Db.NewDelete().
		Model(data).
		Where("user_id = ?", userID).
		Where("action_id = ?", actionID).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func initDB() *bun.DB {
	dsn, err := utils.GetEnvParameter("DATABASE_URL")
	if err != nil {
		log.Println("Err while loading env file: ", dsn)
		return nil
	}
	hsqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	conn := bun.NewDB(hsqldb, pgdialect.New()) // Be careful to create a new db instance each time
	return conn
}
