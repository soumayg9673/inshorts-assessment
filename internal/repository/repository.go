package repository

import (
	"github.com/soumayg9673/inshorts-assessment/internal/database"
	v1rpo "github.com/soumayg9673/inshorts-assessment/internal/repository/v1"
	"go.uber.org/zap"
)

type RepoStore struct {
	Logger *zap.Logger
	ENV    string
}

type Repository struct {
	V1 v1rpo.V1
}

func NewRpoStore(db database.Database, log *zap.Logger, env string) Repository {
	return Repository{
		V1: &v1rpo.V1Rpo{
			DB:  db.V1,
			LOG: log,
			ENV: env,
		},
	}
}
