package dtos

import "gobooru/internal/models"

type ListTagCategoryDTO struct {
}

type ListTagCategoryResponseDTO struct {
	TagCategories []models.TagCategory `json:"tag_categories"`
}
