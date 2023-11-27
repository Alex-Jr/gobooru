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
