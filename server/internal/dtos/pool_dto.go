package dtos

import "gobooru/internal/models"

type CreatePoolDTO struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	PostIDs     []int    `json:"posts"`
	Custom      []string `json:"custom"`
}

type CreatePoolResponseDTO struct {
	Pool models.Pool `json:"pool"`
}
