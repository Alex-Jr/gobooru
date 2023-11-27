package dtos

import "gobooru/internal/models"

type FetchTagDTO struct {
	ID string `param:"id"`
}

type FetchTagResponseDTO struct {
	Tag models.Tag `json:"tag"`
}
