package v1svc

import (
	"github.com/soumayg9673/inshorts-assessment/internal/llm"
	newsv1 "github.com/soumayg9673/inshorts-assessment/internal/models/v1/news"
	v1rpo "github.com/soumayg9673/inshorts-assessment/internal/repository/v1"
	"go.uber.org/zap"
)

type V1Svc struct {
	RPO v1rpo.V1
	LOG *zap.Logger
	Llm llm.Llm
	ENV string
}

type V1 interface {
	GetNewsByCategory([]string) ([]newsv1.NewsApi, error)
	GetNewsByScore() ([]newsv1.NewsApi, error)
	GetNewsBySearch()
	GetNewsBySource(int) ([]newsv1.NewsApi, error)
	GetNewsByNearby()
}
