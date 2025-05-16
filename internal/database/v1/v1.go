package v1db

import (
	"database/sql"

	"go.uber.org/zap"
)

type V1Db struct {
	DB  *sql.DB
	LOG *zap.Logger
	ENV string
}

type V1 interface {
	GetNewsByCategory([]string) (*sql.Rows, error)
	GetNewsByScore()
	GetNewsBySearch()
	GetNewsBySource()
	GetNewsByNearby()
	PatchLlmSummary(string, string) *sql.Row
}
