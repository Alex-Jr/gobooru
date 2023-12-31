package services

import (
	"context"
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type PoolService interface {
	Create(ctx context.Context, dto dtos.CreatePoolDTO) (dtos.CreatePoolResponseDTO, error)
	Delete(ctx context.Context, dto dtos.DeletePoolDTO) (dtos.DeletePoolResponseDTO, error)
	Fetch(ctx context.Context, dto dtos.FetchPoolDTO) (dtos.FetchPoolResponseDTO, error)
	List(ctx context.Context, dto dtos.ListPoolDTO) (dtos.ListPoolResponseDTO, error)
	Update(ctx context.Context, dto dtos.UpdatePoolDTO) (dtos.UpdatePoolResponseDTO, error)
}

type poolService struct {
	poolRepository repositories.PoolRepository
}

type PoolServiceConfig struct {
	PoolRepository repositories.PoolRepository
}

func NewPoolService(c PoolServiceConfig) PoolService {
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
		return dtos.CreatePoolResponseDTO{}, fmt.Errorf("poolRepository.Create: %w", err)
	}

	return dtos.CreatePoolResponseDTO{
		Pool: pool,
	}, nil
}

func (s poolService) Delete(ctx context.Context, dto dtos.DeletePoolDTO) (dtos.DeletePoolResponseDTO, error) {
	pool, err := s.poolRepository.Delete(ctx, dto.ID)
	if err != nil {
		return dtos.DeletePoolResponseDTO{}, fmt.Errorf("poolRepository.Delete: %w", err)
	}

	return dtos.DeletePoolResponseDTO{
		Pool: pool,
	}, nil
}

func (s poolService) Fetch(ctx context.Context, dto dtos.FetchPoolDTO) (dtos.FetchPoolResponseDTO, error) {
	pool, err := s.poolRepository.GetFull(ctx, dto.ID)
	if err != nil {
		return dtos.FetchPoolResponseDTO{}, fmt.Errorf("poolRepository.GetFull: %w", err)
	}

	return dtos.FetchPoolResponseDTO{
		Pool: pool,
	}, nil
}

func (s poolService) List(ctx context.Context, dto dtos.ListPoolDTO) (dtos.ListPoolResponseDTO, error) {
	pools, count, err := s.poolRepository.ListFull(ctx, repositories.PoolListFullArgs{
		Text:     dto.Search,
		Page:     dto.Page,
		PageSize: dto.PageSize,
	})
	if err != nil {
		return dtos.ListPoolResponseDTO{}, fmt.Errorf("poolRepository.ListFull: %w", err)
	}

	return dtos.ListPoolResponseDTO{
		Pools: pools,
		Count: count,
	}, nil
}

func (s poolService) Update(ctx context.Context, dto dtos.UpdatePoolDTO) (dtos.UpdatePoolResponseDTO, error) {
	pool, err := s.poolRepository.Update(ctx, repositories.PoolUpdateArgs{
		Custom:      dto.Custom,
		Description: dto.Description,
		ID:          dto.ID,
		Name:        dto.Name,
		Posts:       dto.PostIDs,
	})
	if err != nil {
		return dtos.UpdatePoolResponseDTO{}, fmt.Errorf("poolRepository.Update: %w", err)
	}

	return dtos.UpdatePoolResponseDTO{
		Pool: pool,
	}, nil
}
