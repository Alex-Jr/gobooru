package repositories

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
)

type TagCategoryRepository interface {
	List(ctx context.Context) ([]models.TagCategory, error)
}

type tagCategoryRepository struct {
	sqlClient        database.SQLClient
	tagCategoryQuery queries.TagCategoryQuery
}

type TagCategoryRepositoryConfig struct {
	SQLClient database.SQLClient
}

func NewTagCategoryRepository(c TagCategoryRepositoryConfig) TagCategoryRepository {
	return &tagCategoryRepository{
		sqlClient:        c.SQLClient,
		tagCategoryQuery: queries.NewTagCategoryQuery(),
	}
}

func (r *tagCategoryRepository) List(ctx context.Context) ([]models.TagCategory, error) {
	var tagCategories []models.TagCategory

	err := r.tagCategoryQuery.List(ctx, r.sqlClient, &tagCategories)
	if err != nil {
		return nil, err
	}

	return tagCategories, nil
}
