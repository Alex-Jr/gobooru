package services

import (
	"context"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type TagCategoryService interface {
	List(ctx context.Context, dtos dtos.ListTagCategoryDTO) (dtos.ListTagCategoryResponseDTO, error)
}

type tagCategoryService struct {
	tagCategoryRepository repositories.TagCategoryRepository
}

type TagCategoryServiceConfig struct {
	TagCategoryRepository repositories.TagCategoryRepository
}

func NewTagCategoryService(c TagCategoryServiceConfig) TagCategoryService {
	return &tagCategoryService{
		tagCategoryRepository: c.TagCategoryRepository,
	}
}

func (s *tagCategoryService) List(ctx context.Context, dto dtos.ListTagCategoryDTO) (dtos.ListTagCategoryResponseDTO, error) {
	tagCategories, err := s.tagCategoryRepository.List(ctx)
	if err != nil {
		return dtos.ListTagCategoryResponseDTO{}, err
	}

	return dtos.ListTagCategoryResponseDTO{
		TagCategories: tagCategories,
	}, nil
}
