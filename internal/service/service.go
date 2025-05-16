package service

import (
	"github.com/soumayg9673/inshorts-assessment/internal/llm"
	"github.com/soumayg9673/inshorts-assessment/internal/repository"
	v1svc "github.com/soumayg9673/inshorts-assessment/internal/service/v1"
	"go.uber.org/zap"
)

type Service struct {
	V1 v1svc.V1
}

func NewServiceStore(rpo repository.Repository, log *zap.Logger, llm llm.Llm, env string) Service {
	return Service{
		V1: &v1svc.V1Svc{
			RPO: rpo.V1,
			LOG: log,
			ENV: env,
		},
	}
}
