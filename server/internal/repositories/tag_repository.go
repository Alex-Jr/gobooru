package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
)

type TagRepository interface {
	Get(ctx context.Context, tagID string) (models.Tag, error)
}

type tagRepository struct {
	sqlClient database.SQLClient
	tagQuery  queries.TagQuery
}

type TagRepositoryConfig struct {
	SQLClient database.SQLClient
}

func NewTagRepository(c TagRepositoryConfig) TagRepository {
	return &tagRepository{
		sqlClient: c.SQLClient,
		tagQuery:  queries.NewTagQuery(),
	}
}

func (r *tagRepository) Get(ctx context.Context, tagID string) (models.Tag, error) {
	tag := models.Tag{ID: tagID}

	err := r.tagQuery.Get(ctx, r.sqlClient, &tag)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tagQuery.Get: %w", err)
	}

	return tag, nil
}
