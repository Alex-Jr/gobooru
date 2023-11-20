package fakes

import (
	"gobooru/internal/models"
	"time"
)

var Post1 = models.Post{
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "post 1 description",
	ID:          1,
	UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Post2 = models.Post{
	CreatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
	Description: "post 2 description",
	ID:          2,
	UpdatedAt:   time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
}

var Post3 = models.Post{
	CreatedAt:   time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
	Description: "post 3 description",
	ID:          3,
	UpdatedAt:   time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
}

var Post4 = models.Post{
	CreatedAt:   time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC),
	Description: "post 4 description",
	ID:          4,
	UpdatedAt:   time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC),
}

var Post5 = models.Post{
	CreatedAt:   time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
	Description: "post 5 description",
	ID:          5,
	UpdatedAt:   time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
}

var Posts = []models.Post{
	Post1,
	Post2,
	Post3,
	Post4,
	Post5,
}

func LoadPosts(p models.Post) models.Post {
	for i := range p.Pools {
		p.Pools[i].Posts = Pools[p.Pools[i].ID-1].Posts
	}

	return p
}
