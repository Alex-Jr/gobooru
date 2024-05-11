package services

import (
	"context"
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type TagService interface {
	Fetch(ctx context.Context, dto dtos.FetchTagDTO) (dtos.FetchTagResponseDTO, error)
	Delete(ctx context.Context, dto dtos.DeleteTagDTO) (dtos.DeleteTagResponseDTO, error)
	List(ctx context.Context, dto dtos.ListTagDTO) (dtos.ListTagResponseDTO, error)
	Update(ctx context.Context, dto dtos.UpdateTagDTO) (dtos.UpdateTagResponseDTO, error)
}

type tagService struct {
	tagRepository repositories.TagRepository
}

type TagServiceConfig struct {
	TagRepository repositories.TagRepository
}

func NewTagService(c TagServiceConfig) TagService {
	return &tagService{
		tagRepository: c.TagRepository,
	}
}

func (s *tagService) Fetch(ctx context.Context, dto dtos.FetchTagDTO) (dtos.FetchTagResponseDTO, error) {
	tag, err := s.tagRepository.Get(ctx, dto.ID)
	if err != nil {
		return dtos.FetchTagResponseDTO{}, fmt.Errorf("tagRepository.Get: %w", err)
	}

	return dtos.FetchTagResponseDTO{
		Tag: tag,
	}, nil
}

func (s *tagService) Delete(ctx context.Context, dto dtos.DeleteTagDTO) (dtos.DeleteTagResponseDTO, error) {
	tag, err := s.tagRepository.Delete(ctx, dto.ID)
	if err != nil {
		return dtos.DeleteTagResponseDTO{}, fmt.Errorf("tagRepository.Delete: %w", err)
	}

	return dtos.DeleteTagResponseDTO{
		Tag: tag,
	}, nil
}

func (s *tagService) List(ctx context.Context, dto dtos.ListTagDTO) (dtos.ListTagResponseDTO, error) {
	tags, count, err := s.tagRepository.List(ctx, repositories.ListTagArgs{
		Search:   dto.Search,
		Page:     dto.Page,
		PageSize: dto.PageSize,
	})
	if err != nil {
		return dtos.ListTagResponseDTO{}, fmt.Errorf("tagRepository.List: %w", err)
	}

	return dtos.ListTagResponseDTO{
		Tags:  tags,
		Count: count,
	}, nil
}

func (s *tagService) Update(ctx context.Context, dto dtos.UpdateTagDTO) (dtos.UpdateTagResponseDTO, error) {
	tag, err := s.tagRepository.Update(ctx, repositories.UpdateTagArgs{
		ID:          dto.ID,
		Description: dto.Description,
		CategoryID:  dto.CategoryID,
	})
	if err != nil {
		return dtos.UpdateTagResponseDTO{}, fmt.Errorf("tagRepository.Update: %w", err)
	}

	return dtos.UpdateTagResponseDTO{
		Tag: tag,
	}, nil
}
