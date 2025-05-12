package database

import (
	"database/sql"

	"go.uber.org/zap"
)

type DbStore struct {
	DB     *sql.DB
	Logger *zap.Logger
	ENV    string
}

type Database struct {
	Database interface {
		QueryRandom()
	}
}

func NewDatabase(db *sql.DB, log *zap.Logger, env string) Database {
	return Database{
		Database: &DbStore{
			DB:     db,
			Logger: log,
			ENV:    env,
		},
	}
}
