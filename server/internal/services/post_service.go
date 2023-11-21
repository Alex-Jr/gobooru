package services

import (
	"context"
	"fmt"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type PostService interface {
	Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error)
	Delete(ctx context.Context, dto dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error)
	Fetch(ctx context.Context, dto dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error)
	List(ctx context.Context, dto dtos.ListPostDTO) (dtos.ListPostResponseDTO, error)
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
		Rating:      dto.Rating,
		Tags:        dto.Tags,
	})

	if err != nil {
		return dtos.CreatePostResponseDTO{}, fmt.Errorf("postRepository.Create: %w", err)
	}

	return dtos.CreatePostResponseDTO{
		Post: post,
	}, nil
}

func (s postService) Delete(ctx context.Context, dto dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error) {
	post, err := s.postRepository.GetFull(ctx, dto.ID)
	if err != nil {
		return dtos.DeletePostResponseDTO{}, fmt.Errorf("postRepository.GetFull: %w", err)
	}

	err = s.postRepository.Delete(ctx, dto.ID)
	if err != nil {
		return dtos.DeletePostResponseDTO{}, fmt.Errorf("postRepository.Delete: %w", err)
	}

	return dtos.DeletePostResponseDTO{
		Post: post,
	}, nil
}

func (s postService) Fetch(ctx context.Context, dto dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error) {
	post, err := s.postRepository.GetFull(ctx, dto.ID)
	if err != nil {
		return dtos.FetchPostResponseDTO{}, fmt.Errorf("postRepository.GetFull: %w", err)
	}

	return dtos.FetchPostResponseDTO{
		Post: post,
	}, nil
}

func (s postService) List(ctx context.Context, dto dtos.ListPostDTO) (dtos.ListPostResponseDTO, error) {
	posts, count, err := s.postRepository.List(ctx, repositories.ListPostsArgs{
		Search:   dto.Search,
		Page:     dto.Page,
		PageSize: dto.PageSize,
	})

	if err != nil {
		return dtos.ListPostResponseDTO{}, fmt.Errorf("postRepository.List: %w", err)
	}

	return dtos.ListPostResponseDTO{
		Posts: posts,
		Count: count,
	}, nil
}
