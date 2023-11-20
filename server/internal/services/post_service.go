package services

import (
	"context"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type PostService interface {
	Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error)
	Delete(ctx context.Context, dto dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error)
	Fetch(ctx context.Context, dto dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error)
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

func (s postService) Delete(ctx context.Context, dto dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error) {
	post, err := s.postRepository.GetFull(ctx, dto.ID)
	if err != nil {
		return dtos.DeletePostResponseDTO{}, err
	}

	err = s.postRepository.Delete(ctx, dto.ID)
	if err != nil {
		return dtos.DeletePostResponseDTO{}, err
	}

	return dtos.DeletePostResponseDTO{
		Post: post,
	}, nil
}

func (s postService) Fetch(ctx context.Context, dto dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error) {
	post, err := s.postRepository.GetFull(ctx, dto.ID)
	if err != nil {
		return dtos.FetchPostResponseDTO{}, err
	}

	return dtos.FetchPostResponseDTO{
		Post: post,
	}, nil
}
