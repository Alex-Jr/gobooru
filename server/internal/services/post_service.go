package services

import (
	"context"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type PostService interface {
	Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error)
}

type postService struct {
	postRepository repositories.PostRepository
}

type PostServiceConfig struct {
	PostRepository repositories.PostRepository
}

func NewPostService(c PostServiceConfig) PostService {
	return &postService{
		postRepository: c.PostRepository,
	}
}

func (s postService) Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error) {
	post, err := s.postRepository.Create(ctx, repositories.CreatePostArgs{
		Description: dto.Description,
	})

	if err != nil {
		return dtos.CreatePostResponseDTO{}, err
	}

	return dtos.CreatePostResponseDTO{
		Post: post,
	}, nil
}
