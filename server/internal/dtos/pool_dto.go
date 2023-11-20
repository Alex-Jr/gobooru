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

type DeletePoolDTO struct {
	ID int `param:"id"`
}

type DeletePoolResponseDTO struct {
	Pool models.Pool `json:"pool"`
}

type FetchPoolDTO struct {
	ID int `param:"id"`
}

type FetchPoolResponseDTO struct {
	Pool models.Pool `json:"pool"`
}

type ListPoolDTO struct {
	Search   string `query:"search"`
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
}

type ListPoolResponseDTO struct {
	Pools []models.Pool `json:"pools"`
	Count int           `json:"count"`
}
