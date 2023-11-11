package dtos

import "gobooru/internal/models"

type CreatePoolDTO struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	PostIDs     []int  `json:"posts"`
	UserID      int
}

type CreatePoolResponseDTO struct {
	Pool models.Pool `json:"pool"`
}
