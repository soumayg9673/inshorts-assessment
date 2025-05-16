package v1rpo

import (
	v1db "github.com/soumayg9673/inshorts-assessment/internal/database/v1"
	newsv1 "github.com/soumayg9673/inshorts-assessment/internal/models/v1/news"
	"go.uber.org/zap"
)

type V1Rpo struct {
	DB  v1db.V1
	LOG *zap.Logger
	ENV string
}

type V1 interface {
	GetNewsByCategory([]string) ([]newsv1.NewsSql, error)
	GetNewsByScore() ([]newsv1.NewsSql, error)
	GetNewsBySearch()
	GetNewsBySource()
	GetNewsByNearby()
	PatchLlmSummary(string, string)
}
