package service

import (
	"JMIND/internal/domain"
	"JMIND/internal/query"
	"JMIND/internal/repository"
	"context"
)

type Eth interface {
	GetInfoAboutBlock(ctx context.Context, number query.BlockByNumber) (domain.BlockInfo, error)
}
type Service struct {
	Eth
}

func NewService(repository *repository.Repository) *Service {
	return &Service{Eth: NewEthService(repository)}
}
