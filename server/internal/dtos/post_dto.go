package dtos

import "gobooru/internal/models"

type CreatePostDTO struct {
	Description string `form:"description"`
}

type CreatePostResponseDTO struct {
	Post models.Post `json:"post"`
}
