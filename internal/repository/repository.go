package repository

import "go.uber.org/zap"

type RepoStore struct {
	Logger *zap.Logger
	ENV    string
}
