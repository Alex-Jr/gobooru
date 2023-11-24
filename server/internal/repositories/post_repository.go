package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
	"gobooru/internal/slice_utils"
	"time"
)

type PostRepository interface {
	Create(ctx context.Context, args CreatePostArgs) (models.Post, error)
	Delete(ctx context.Context, postID int) error
	GetFull(ctx context.Context, postID int) (models.Post, error)
	List(ctx context.Context, args ListPostsArgs) ([]models.Post, int, error)
	Update(ctx context.Context, args UpdatePostArgs) (models.Post, error)
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

	tagsDeduped := slice_utils.Deduplicate(args.Tags)

	tags := make([]models.Tag, len(tagsDeduped))

	post := models.Post{
		Rating:      args.Rating,
		Description: args.Description,
		TagIDs:      make([]string, len(tags)),
		TagCount:    len(tags),
		PoolCount:   0,
		CreatedAt:   now,
		UpdatedAt:   now,
		Pools:       make(models.PoolList, 0),
		Tags:        make(models.TagList, len(tags)),
	}

	err = r.postQuery.Create(ctx, tx, &post)

	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.Create: %w", err)
	}

	for i, tag := range tagsDeduped {
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
		return fmt.Errorf("postQuery.Delete: %w", err)
	}

	return nil
}

func (r *postRepository) GetFull(ctx context.Context, postID int) (models.Post, error) {
	post := models.Post{
		ID: postID,
	}

	err := r.postQuery.GetFull(ctx, r.sqlClient, &post)

	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFull: %w", err)
	}

	return post, nil
}

type ListPostsArgs struct {
	Search   string
	Page     int
	PageSize int
}

func (r *postRepository) List(ctx context.Context, args ListPostsArgs) ([]models.Post, int, error) {
	posts := make([]models.Post, 0)
	count := 0

	err := r.postQuery.List(
		ctx,
		r.sqlClient,
		models.Search{
			Text:     args.Search,
			Page:     args.Page,
			PageSize: args.PageSize,
		},
		&posts,
		&count,
	)

	if err != nil {
		return nil, 0, fmt.Errorf("postQuery.List: %w", err)
	}

	return posts, count, nil
}

type UpdatePostArgs struct {
	ID          int
	Description *string
	Rating      *string
	Tags        *[]string
}

func (r *postRepository) Update(ctx context.Context, args UpdatePostArgs) (models.Post, error) {
	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Post{}, fmt.Errorf("sqlClient.BeginTxx: %w", err)
	}

	post := models.Post{
		ID: args.ID,
	}

	err = r.postQuery.GetFull(ctx, tx, &post)
	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFull: %w", err)
	}

	if args.Description != nil {
		post.Description = *args.Description
	}

	if args.Rating != nil {
		post.Rating = *args.Rating
	}

	if args.Tags != nil {
		// TODO: resolve implications and alias

		toRemove := slice_utils.Difference(post.TagIDs, *args.Tags)
		toAdd := slice_utils.Difference(*args.Tags, post.TagIDs)

		if len(toRemove) > 0 {
			err = r.postTag.DisassociatePostsByID(ctx, tx, post, toRemove)
			if err != nil {
				return models.Post{}, fmt.Errorf("postTag.DisassociatePosts: %w", err)
			}

			err = r.tagQuery.UpdatePostCount(ctx, tx, toRemove, -1)
			if err != nil {
				return models.Post{}, fmt.Errorf("tagQuery.UpdatePostCount: %w", err)
			}
		}

		if len(toAdd) > 0 {
			tags := make([]models.Tag, len(toAdd))

			now := time.Now()
			for i, tag := range toAdd {
				tags[i] = models.Tag{
					ID:        tag,
					PostCount: 1,
					CreatedAt: now,
					UpdatedAt: now,
				}
			}

			err = r.tagQuery.CreateMany(ctx, tx, &tags)
			if err != nil {
				return models.Post{}, fmt.Errorf("tagQuery.CreateMany: %w", err)
			}

			err = r.postTag.AssociatePosts(ctx, tx, post, tags)
			if err != nil {
				return models.Post{}, fmt.Errorf("postTag.AssociatePosts: %w", err)
			}
		}

		post.TagIDs = *args.Tags

		post.TagCount = len(*args.Tags)
	}

	err = r.postQuery.Update(ctx, tx, post)
	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.Update: %w", err)
	}

	err = r.postQuery.GetFull(ctx, tx, &post)
	if err != nil {
		return models.Post{}, fmt.Errorf("postQuery.GetFull: %w", err)
	}

	tx.Commit()

	return post, nil
}
