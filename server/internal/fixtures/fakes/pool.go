package fakes

import (
	"gobooru/internal/models"
	"time"

	"github.com/lib/pq"
)

var Pool1 = models.Pool{
	Custom:      pq.StringArray{"a"},
	CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "pool 1 description",
	ID:          1,
	Name:        "pool 1 name",
	PostCount:   1,
	Posts: []models.Post{
		{
			ID: 1,
		},
	},
	PostIDs:   pq.Int64Array{1},
	UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Pool2 = models.Pool{
	Custom:      pq.StringArray{},
	CreatedAt:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "pool 2 description",
	ID:          2,
	Name:        "pool 2 name",
	PostCount:   2,
	Posts: []models.Post{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	},
	PostIDs:   pq.Int64Array{1, 2},
	UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Pool3 = models.Pool{
	Custom:      pq.StringArray{},
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "pool 3 description",
	ID:          3,
	Name:        "pool 3 name",
	PostCount:   3,
	Posts: []models.Post{
		{
			ID: 3,
		},
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	},
	PostIDs:   pq.Int64Array{3, 1, 2},
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Pool4 = models.Pool{
	Custom:      pq.StringArray{"shared"},
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "pool 4 description",
	ID:          4,
	Name:        "pool 4 name",
	PostCount:   1,
	Posts: []models.Post{
		{
			ID: 1,
		},
	},
	PostIDs:   pq.Int64Array{1},
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Pool5 = models.Pool{
	Custom:      pq.StringArray{"shared"},
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "pool 5 description",
	ID:          5,
	Name:        "pool 5 name",
	PostCount:   1,
	Posts: []models.Post{
		{
			ID: 2,
		},
	},
	PostIDs:   pq.Int64Array{2},
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Pool6 = models.Pool{
	Custom:      pq.StringArray{"shared"},
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "pool 6 description",
	ID:          6,
	Name:        "pool 6 name",
	PostCount:   1,
	Posts: []models.Post{
		{
			ID: 3,
		},
	},
	PostIDs:   pq.Int64Array{3},
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Pools = []models.Pool{
	Pool1,
	Pool2,
	Pool3,
	Pool4,
	Pool5,
	Pool6,
}

// function to load relationships without cyclic dependencies
func LoadPoolRelations(p models.Pool) models.Pool {
	for i := range p.Posts {
		p.Posts[i] = Posts[p.Posts[i].ID-1]
		p.Posts[i].Pools = nil
		p.Posts[i].Tags = nil
		p.Posts[i].Relations = nil
	}

	return p
}
