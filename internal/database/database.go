package database

import (
	"database/sql"

	v1db "github.com/soumayg9673/inshorts-assessment/internal/database/v1"
	"go.uber.org/zap"
)

type DbStore struct {
	DB  *sql.DB
	LOG *zap.Logger
	ENV string
}

type Database struct {
	Database interface {
		QueryRandom()
	}
	V1 v1db.V1
}

func NewDbStore(db *sql.DB, log *zap.Logger, env string) Database {
	return Database{
		Database: &DbStore{
			DB:  db,
			LOG: log,
			ENV: env,
		},
		V1: &v1db.V1Db{
			DB:  db,
			LOG: log,
			ENV: env,
		},
	}
}
