package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
	"time"
)

type TagRepository interface {
	Delete(ctx context.Context, tagID string) (models.Tag, error)
	Get(ctx context.Context, tagID string) (models.Tag, error)
	List(ctx context.Context, args ListTagArgs) ([]models.Tag, int, error)
	Update(ctx context.Context, args UpdateTagArgs) (models.Tag, error)
}

type tagRepository struct {
	sqlClient   database.SQLClient
	postQuery   queries.PostQuery
	tagQuery    queries.TagQuery
	tagCategory queries.TagCategoryQuery
}

type TagRepositoryConfig struct {
	SQLClient database.SQLClient
}

func NewTagRepository(c TagRepositoryConfig) TagRepository {
	return &tagRepository{
		sqlClient:   c.SQLClient,
		tagQuery:    queries.NewTagQuery(),
		postQuery:   queries.NewPostQuery(),
		tagCategory: queries.NewTagCategoryQuery(),
	}
}

func (r *tagRepository) Delete(ctx context.Context, tagID string) (models.Tag, error) {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Tag{}, fmt.Errorf("sqlClient.BeginTx: %w", err)
	}
	defer tx.Rollback()

	tag := models.Tag{ID: tagID}

	err = r.tagQuery.Get(ctx, tx, &tag)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tagQuery.Get: %w", err)
	}

	err = r.tagQuery.Delete(ctx, tx, tag)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tagQuery.Delete: %w", err)
	}

	err = r.postQuery.RemoveTag(ctx, tx, tag.ID)
	if err != nil {
		return models.Tag{}, fmt.Errorf("postQuery.RemoveTag: %w", err)
	}

	err = r.tagCategory.UpdateTagCount(ctx, tx, tag.CategoryID, -1)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tagCategory.UpdateTagCount: %w", err)
	}

	tx.Commit()

	return tag, nil
}

func (r *tagRepository) Get(ctx context.Context, tagID string) (models.Tag, error) {
	tag := models.Tag{ID: tagID}

	err := r.tagQuery.Get(ctx, r.sqlClient, &tag)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tagQuery.Get: %w", err)
	}

	return tag, nil
}

type ListTagArgs struct {
	Search   string
	Page     int
	PageSize int
}

func (r *tagRepository) List(ctx context.Context, args ListTagArgs) ([]models.Tag, int, error) {
	tags := []models.Tag{}
	count := 0

	err := r.tagQuery.List(
		ctx,
		r.sqlClient,
		models.Search{
			Text:     args.Search,
			Page:     args.Page,
			PageSize: args.PageSize,
		},
		&count,
		&tags,
	)
	if err != nil {
		return tags, count, fmt.Errorf("tagQuery.List: %w", err)
	}

	return tags, count, nil
}

type UpdateTagArgs struct {
	ID          string
	Description *string
	CategoryID  *string
}

func (r *tagRepository) Update(ctx context.Context, args UpdateTagArgs) (models.Tag, error) {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Tag{}, fmt.Errorf("sqlClient.BeginTx: %w", err)
	}
	defer tx.Rollback()

	tag := models.Tag{ID: args.ID}

	err = r.tagQuery.Get(ctx, tx, &tag)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tagQuery.Get: %w", err)
	}

	if args.Description != nil {
		tag.Description = *args.Description
	}

	if args.CategoryID != nil && *args.CategoryID != tag.CategoryID {
		err = r.tagCategory.UpdateTagCount(ctx, tx, *args.CategoryID, 1)
		if err != nil {
			return models.Tag{}, fmt.Errorf("tagCategory.UpdateTagCount: %w", err)
		}

		err = r.tagCategory.UpdateTagCount(ctx, tx, tag.CategoryID, -1)
		if err != nil {
			return models.Tag{}, fmt.Errorf("tagCategory.UpdateTagCount: %w", err)
		}

		tag.CategoryID = *args.CategoryID
	}

	tag.UpdatedAt = time.Now()

	err = r.tagQuery.Update(ctx, tx, tag)
	if err != nil {
		return models.Tag{}, fmt.Errorf("tag.Update: %w", err)
	}

	tx.Commit()

	return tag, nil
}
