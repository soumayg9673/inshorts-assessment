package v1svc

import (
	newsv1 "github.com/soumayg9673/inshorts-assessment/internal/models/v1/news"
	v1rpo "github.com/soumayg9673/inshorts-assessment/internal/repository/v1"
	"go.uber.org/zap"
)

type V1Svc struct {
	RPO v1rpo.V1
	LOG *zap.Logger
	ENV string
}

type V1 interface {
	GetNewsByCategory([]string) ([]newsv1.NewsApi, error)
	GetNewsByScore()
	GetNewsBySearch()
	GetNewsBySource()
	GetNewsByNearby()
}
