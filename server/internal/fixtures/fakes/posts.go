package fakes

import (
	"gobooru/internal/models"
	"time"
)

var Post1 = models.Post{
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "post 1 description",
	ID:          1,
	PoolCount:   4,
	Rating:      "S",
	TagCount:    1,
	TagIDs:      []string{"tag_one"},
	MD5:         "1",
	FileExt:     "jpg",
	FileSize:    100,
	FilePath:    "1.jpg",
	ThumbPath:   "1-thumb.webp",
	Tags: []models.Tag{
		{
			ID: "tag_one",
		},
	},
	Pools: []models.Pool{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
		{
			ID: 3,
		},
		{
			ID: 4,
		},
	},

	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Post2 = models.Post{
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "post 2 description",
	ID:          2,
	PoolCount:   3,
	Rating:      "S",
	TagCount:    0,
	TagIDs:      []string{},
	MD5:         "2",
	FileExt:     "jpg",
	FileSize:    100,
	FilePath:    "2.jpg",
	ThumbPath:   "2-thumb.webp",
	Tags:        nil,
	Pools: []models.Pool{
		{
			ID: 2,
		},
		{
			ID: 3,
		},
		{
			ID: 5,
		},
	},
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Post3 = models.Post{
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "post 3 description",
	ID:          3,
	PoolCount:   2,
	Rating:      "S",
	TagCount:    0,
	TagIDs:      []string{},
	Tags:        nil,
	MD5:         "3",
	FileExt:     "jpg",
	FileSize:    100,
	FilePath:    "3.jpg",
	ThumbPath:   "3-thumb.webp",
	Pools: []models.Pool{
		{
			ID: 3,
		},
		{
			ID: 6,
		},
	},
	UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Post4 = models.Post{
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "post 4 description",
	ID:          4,
	PoolCount:   0,
	Rating:      "Q",
	TagCount:    0,
	TagIDs:      []string{},
	MD5:         "4",
	FileExt:     "jpg",
	FileSize:    100,
	FilePath:    "4.jpg",
	ThumbPath:   "4-thumb.webp",
	UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Post5 = models.Post{
	CreatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	Description: "post 5 description",
	ID:          5,
	PoolCount:   0,
	Rating:      "E",
	TagCount:    0,
	TagIDs:      []string{},
	MD5:         "5",
	FileExt:     "png",
	FileSize:    1000,
	FilePath:    "5.png",
	ThumbPath:   "5-thumb.webp",
	UpdatedAt:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
}

var Posts = []models.Post{
	Post1,
	Post2,
	Post3,
	Post4,
	Post5,
}

func LoadPostRelations(p models.Post) models.Post {
	for i := range p.Pools {
		p.Pools[i] = Pools[p.Pools[i].ID-1]
		p.Pools[i].Posts = nil
	}

	for i := range p.Tags {
		p.Tags[i] = Tags[p.Tags[i].ID]
	}

	return p
}

func LoadPostNoRelations(p models.Post) models.Post {
	p.Pools = nil
	p.Tags = nil

	return p
}
