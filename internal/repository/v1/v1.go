package v1rpo

import (
	v1db "github.com/soumayg9673/inshorts-assessment/internal/database/v1"
	"go.uber.org/zap"
)

type V1Rpo struct {
	DB  v1db.V1Db
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
