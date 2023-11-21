package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
	"time"
)

type PostRepository interface {
	Create(ctx context.Context, args CreatePostArgs) (models.Post, error)
	Delete(ctx context.Context, postID int) error
	GetFull(ctx context.Context, postID int) (models.Post, error)
}

type postRepository struct {
	sqlClient database.SQLClient
	postQuery queries.PostQuery
	tagQuery  queries.TagQuery
	postTag   queries.PostTagQuery
}

type CreatePostArgs struct {
	Description string
	Rating      string
	Tags        []string
}

func NewPostRepository(sqlClient database.SQLClient) PostRepository {
	return &postRepository{
		sqlClient: sqlClient,
		postQuery: queries.NewPostQuery(),
		tagQuery:  queries.NewTagQuery(),
		postTag:   queries.NewPostTagQuery(),
	}
}

func (r *postRepository) Create(ctx context.Context, args CreatePostArgs) (models.Post, error) {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Post{}, fmt.Errorf("sqlClient.BeginTxx: %w", err)
	}
	defer tx.Rollback()

	now := time.Now()

	post := models.Post{
		Rating:      args.Rating,
		Description: args.Description,
		TagIDs:      make([]string, len(args.Tags)),
		TagCount:    len(args.Tags),
		PoolCount:   0,
		CreatedAt:   now,
		UpdatedAt:   now,
		Pools:       make(models.PoolList, 0),
		Tags:        make(models.TagList, len(args.Tags)),
	}

	err = r.postQuery.Create(ctx, tx, &post)

	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.Create: %w", err)
	}

	tags := make([]models.Tag, len(args.Tags))

	for i, tag := range args.Tags {
		tags[i] = models.Tag{
			ID:        tag,
			PostCount: 1,
			CreatedAt: now,
			UpdatedAt: now,
		}

		post.TagIDs[i] = tag
	}
	post.Tags = tags

	err = r.tagQuery.CreateMany(ctx, tx, &tags)

	if err != nil {
		return models.Post{}, fmt.Errorf("tagQuery.CreateMany: %w", err)
	}

	err = r.postTag.AssociatePosts(ctx, tx, post, tags)
	if err != nil {
		return models.Post{}, fmt.Errorf("postTag.AssociatePosts: %w", err)
	}

	tx.Commit()

	return post, nil
}

func (r *postRepository) Delete(ctx context.Context, postID int) error {
	post := models.Post{
		ID: postID,
	}

	err := r.postQuery.Delete(ctx, r.sqlClient, &post)

	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) GetFull(ctx context.Context, postID int) (models.Post, error) {
	post := models.Post{
		ID: postID,
	}

	err := r.postQuery.GetFull(ctx, r.sqlClient, &post)

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}
