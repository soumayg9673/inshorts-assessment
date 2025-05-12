package service

import (
	"github.com/soumayg9673/inshorts-assessment/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
}

func NewServiceStore(rpo repository.Repository, log *zap.Logger, env string) Service {
	return Service{}
}
