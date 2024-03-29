package services

import (
	"context"
	"errors"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/dtos"
	"gobooru/internal/repositories"
)

type PostService interface {
	Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error)
	Delete(ctx context.Context, dto dtos.DeletePostDTO) (dtos.DeletePostResponseDTO, error)
	Fetch(ctx context.Context, dto dtos.FetchPostDTO) (dtos.FetchPostResponseDTO, error)
	FetchByHash(ctx context.Context, dto dtos.FetchPostByHashDTO) (dtos.FetchPostByHashResponseDTO, error)
	List(ctx context.Context, dto dtos.ListPostDTO) (dtos.ListPostResponseDTO, error)
	Update(ctx context.Context, dto dtos.UpdatePostDTO) (dtos.UpdatePostResponseDTO, error)
	CreateNote(ctx context.Context, dto dtos.CreatePostNoteDTO) (dtos.CreatePostNoteResponseDTO, error)
}

type postService struct {
	postRepository      repositories.PostRepository
	postNotesRepository repositories.PostNotesRepository
	fileService         FileService
	IQDBService         IQDBService
}

type PostServiceConfig struct {
	PostRepository      repositories.PostRepository
	PostNotesRepository repositories.PostNotesRepository
	FileService         FileService
	IQDBService         IQDBService
}

func NewPostService(c PostServiceConfig) PostService {
	return &postService{
		postRepository:      c.PostRepository,
		postNotesRepository: c.PostNotesRepository,
		fileService:         c.FileService,
		IQDBService:         c.IQDBService,
	}
}

func (s postService) Create(ctx context.Context, dto dtos.CreatePostDTO) (dtos.CreatePostResponseDTO, error) {
	file, err := s.fileService.HandleUpload(dto.File)
	if err != nil {
		return dtos.CreatePostResponseDTO{}, fmt.Errorf("fileService.HandleFileUpload: %w", err)
	}

	post, err := s.postRepository.GetFullByHash(ctx, file.MD5)
	if err != nil {
		if !errors.Is(err, database.ErrNotFound) {
			return dtos.CreatePostResponseDTO{}, fmt.Errorf("postRepository.GetByHash: %w", err)
		}
	}

	if post.ID != 0 {
		return dtos.CreatePostResponseDTO{
			Post: post,
		}, nil
	}

	post, err = s.postRepository.Create(ctx, repositories.CreatePostArgs{
		Custom:           dto.Custom,
		Description:      dto.Description,
		Rating:           dto.Rating,
		Tags:             dto.Tags,
		FileExt:          file.FileExt,
		FileSize:         file.FileSize,
		FilePath:         file.FilePath,
		FileOriginalName: file.FileOriginalName,
		ThumbPath:        file.ThumbPath,
		MD5:              file.MD5,
		Sources:          dto.Sources,
		Width:            file.Width,
		Height:           file.Height,
		Duration:         file.Duration,
	})

	if err != nil {
		return dtos.CreatePostResponseDTO{}, fmt.Errorf("postRepository.Create: %w", err)
	}

	// TODO: Handle IQDB errors
	relations, err := s.IQDBService.HandlePost(post)
	if err != nil {
		return dtos.CreatePostResponseDTO{}, fmt.Errorf("iqdbService.HandlePost: %w", err)
	}

	err = s.postRepository.SaveRelations(ctx, &post, &relations)
	if err != nil {
		return dtos.CreatePostResponseDTO{}, fmt.Errorf("postRepository.SaveRelations: %w", err)
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

func (s postService) FetchByHash(ctx context.Context, dto dtos.FetchPostByHashDTO) (dtos.FetchPostByHashResponseDTO, error) {
	post, err := s.postRepository.GetFullByHash(ctx, dto.Hash)
	if err != nil {
		return dtos.FetchPostByHashResponseDTO{}, fmt.Errorf("postRepository.GetByHash: %w", err)
	}

	return dtos.FetchPostByHashResponseDTO{
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

func (s postService) Update(ctx context.Context, dto dtos.UpdatePostDTO) (dtos.UpdatePostResponseDTO, error) {
	post, err := s.postRepository.Update(ctx, repositories.UpdatePostArgs{
		ID:          dto.ID,
		Description: dto.Description,
		Rating:      dto.Rating,
		Tags:        dto.Tags,
		Sources:     dto.Sources,
		Custom:      dto.Custom,
	})

	if err != nil {
		return dtos.UpdatePostResponseDTO{}, fmt.Errorf("postRepository.Update: %w", err)
	}

	return dtos.UpdatePostResponseDTO{
		Post: post,
	}, nil
}

func (s postService) CreateNote(ctx context.Context, dto dtos.CreatePostNoteDTO) (dtos.CreatePostNoteResponseDTO, error) {
	postNote, err := s.postNotesRepository.Create(
		ctx,
		repositories.CreatePostNotesArgs{
			PostID: dto.PostID,
			Body:   dto.Body,
			X:      dto.X,
			Y:      dto.Y,
			Width:  dto.Width,
			Height: dto.Height,
		},
	)

	if err != nil {
		return dtos.CreatePostNoteResponseDTO{}, fmt.Errorf("postNotesRepository.Create: %w", err)
	}

	return dtos.CreatePostNoteResponseDTO{
		PostNote: postNote,
	}, nil
}
