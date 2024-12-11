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
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
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

func initDB() *bun.DB {
	dsn, err := utils.GetEnvParameter("DATABASE_URL")
	if err != nil {
		return nil
	}
	hsqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	conn := bun.NewDB(hsqldb, pgdialect.New()) // Be careful to create a new db instance each time
	return conn
}
