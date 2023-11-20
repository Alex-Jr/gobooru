package fakes

import (
	"gobooru/internal/models"
	"time"
)

var Tag1 = models.Tag{
	ID:          "tag_one",
	Description: "tag one description",
	PostCount:   1,
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Tags = map[string]models.Tag{
	"tag_one": Tag1,
}
