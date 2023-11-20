package dtos

import "gobooru/internal/models"

type CreatePostDTO struct {
	Description string `form:"description"`
	Tags        string `form:"tags"`
}

type CreatePostResponseDTO struct {
	Post models.Post `json:"post"`
}

type DeletePostDTO struct {
	ID int `param:"id"`
}

type DeletePostResponseDTO struct {
	Post models.Post `json:"post"`
}

type FetchPostDTO struct {
	ID int `param:"id"`
}

type FetchPostResponseDTO struct {
	Post models.Post `json:"post"`
}
