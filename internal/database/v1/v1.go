package v1db

import (
	"database/sql"

	"go.uber.org/zap"
)

type DbStoreV1 struct {
	DB  *sql.DB
	LOG *zap.Logger
	ENV string
}

type V1 interface {
	GetNewsByCategory()
	GetNewsByScore()
	GetNewsBySearch()
	GetNewsBySource()
	GetNewsByNearby()
}
