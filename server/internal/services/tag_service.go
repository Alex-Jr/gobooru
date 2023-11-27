package services

import (
	"context"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type TagService interface {
	Fetch(ctx context.Context, dto dtos.FetchTagDTO) (dtos.FetchTagResponseDTO, error)
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
		return dtos.FetchTagResponseDTO{}, err
	}

	return dtos.FetchTagResponseDTO{
		Tag: tag,
	}, nil
}
