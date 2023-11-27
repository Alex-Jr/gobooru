package dtos

import "gobooru/internal/models"

type FetchTagDTO struct {
	ID string `param:"id"`
}

type FetchTagResponseDTO struct {
	Tag models.Tag `json:"tag"`
}

type DeleteTagDTO struct {
	ID string `param:"id"`
}

type DeleteTagResponseDTO struct {
	Tag models.Tag `json:"tag"`
}

type ListTagDTO struct {
	Search   string `query:"search"`
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
}

type ListTagResponseDTO struct {
	Tags  []models.Tag `json:"tags"`
	Count int          `json:"count"`
}
