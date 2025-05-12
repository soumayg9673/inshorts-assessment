package v1svc

import (
	v1rpo "github.com/soumayg9673/inshorts-assessment/internal/repository/v1"
	"go.uber.org/zap"
)

type V1Svc struct {
	RPO v1rpo.V1
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
