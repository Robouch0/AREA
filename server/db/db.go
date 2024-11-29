//
// EPITECH PROJECT, 2024
// AREA
// File description:
// db
//

package db

import (
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	dbConnOnce sync.Once
)

func initDB() *bun.DB {
	dsn := "postgresql://postgres:postgres@localhost:5432/postgres?schema=public" // In environment
	hsqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	conn := bun.NewDB(hsqldb, pgdialect.New()) // Be careful to create a new db instance each time
	return conn
}
