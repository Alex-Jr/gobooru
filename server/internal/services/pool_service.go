package services

import (
	"context"
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type PoolService interface {
	Create(ctx context.Context, dto dtos.CreatePoolDTO) (dtos.CreatePoolResponseDTO, error)
}

type poolService struct {
	poolRepository repositories.PoolRepository
}

type PoolRepositoryConfig struct {
	PoolRepository repositories.PoolRepository
}

func NewPoolService(c PoolRepositoryConfig) PoolService {
	return &poolService{
		poolRepository: c.PoolRepository,
	}
}

func (s poolService) Create(ctx context.Context, dto dtos.CreatePoolDTO) (dtos.CreatePoolResponseDTO, error) {
	pool, err := s.poolRepository.Create(ctx, repositories.PoolCreateArgs{
		Custom:      dto.Custom,
		Description: dto.Description,
		Name:        dto.Name,
		PostIDs:     dto.PostIDs,
	})

	if err != nil {
		return dtos.CreatePoolResponseDTO{}, fmt.Errorf("failed to create pool: %w", err)
	}

	return dtos.CreatePoolResponseDTO{
		Pool: pool,
	}, nil
}
