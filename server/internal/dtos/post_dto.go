package dtos

import (
	"gobooru/internal/models"
	"mime/multipart"
)

type CreatePostDTO struct {
	Custom      []string              `form:"custom"`
	Description string                `form:"description"`
	File        *multipart.FileHeader `form:"file"`
	Rating      string                `form:"rating"`
	Sources     []string              `form:"sources"`
	Tags        []string              `form:"tags"`
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

type FetchPostByHashDTO struct {
	Hash string `param:"hash"`
}

type FetchPostByHashResponseDTO struct {
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
	Sources     *[]string `json:"sources"`
	Custom      *[]string `json:"custom"`
}

type UpdatePostResponseDTO struct {
	Post models.Post `json:"post"`
}
