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
	CategoryID:  "general",
}

var Tag2 = models.Tag{
	ID:          "tag_two",
	Description: "tag two description",
	PostCount:   1,
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	CategoryID:  "general",
}

var Tags = map[string]models.Tag{
	"tag_one": Tag1,
	"tag_two": Tag2,
}

func LoadTagRelations(t models.Tag) models.Tag {
	return t
}
