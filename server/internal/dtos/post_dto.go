package dtos

import "gobooru/internal/models"

type CreatePostDTO struct {
	Description string   `form:"description"`
	Rating      string   `form:"rating"`
	Tags        []string `form:"tags"`
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

type ListPostDTO struct {
	Search   string `query:"search"`
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
}

type ListPostResponseDTO struct {
	Posts []models.Post `json:"posts"`
	Count int           `json:"count"`
}

type UpdatePostDTO struct {
	ID          int       `param:"id"`
	Description *string   `json:"description"`
	Rating      *string   `json:"rating"`
	Tags        *[]string `json:"tags"`
}

type UpdatePostResponseDTO struct {
	Post models.Post `json:"post"`
}
