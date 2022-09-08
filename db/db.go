package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func NewDB() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	DB := bun.NewDB(sqldb, sqlitedialect.New())
	return DB
}

func SetSampleRows(db *bun.DB) error {
	// Create tables
	res, err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx) // Where is ctx from?
}
