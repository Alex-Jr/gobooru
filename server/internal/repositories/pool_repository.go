package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
	"gobooru/internal/slice_utils"
	"time"

	"github.com/lib/pq"
)

type PoolRepository interface {
	Create(ctx context.Context, args PoolCreateArgs) (models.Pool, error)
	Delete(ctx context.Context, id int) (models.Pool, error)
	GetFull(ctx context.Context, id int) (models.Pool, error)
	ListFull(ctx context.Context, args PoolListFullArgs) ([]models.Pool, int, error)
	Update(ctx context.Context, args PoolUpdateArgs) (models.Pool, error)
}

type poolRepository struct {
	sqlClient     database.SQLClient
	poolQuery     queries.PoolQuery
	poolPostQuery queries.PoolPostQuery
	postQuery     queries.PostQuery
}

func NewPoolRepository(dbClient database.SQLClient) PoolRepository {
	return &poolRepository{
		sqlClient:     dbClient,
		poolQuery:     queries.NewPoolQuery(),
		poolPostQuery: queries.NewPoolPostQuery(),
		postQuery:     queries.NewPostQuery(),
	}
}

type PoolCreateArgs struct {
	Custom      []string
	Description string
	Name        string
	PostIDs     []int
}

func (r poolRepository) Create(ctx context.Context, args PoolCreateArgs) (models.Pool, error) {
	// TODO: maybe should use int64 all over the place
	postIDs := make([]int64, len(args.PostIDs))
	for i, id := range args.PostIDs {
		postIDs[i] = int64(id)
	}

	pool := models.Pool{
		Description: args.Description,
		ID:          0,
		Name:        args.Name,
		PostIDs:     postIDs,
		PostCount:   len(args.PostIDs),
		Posts:       make([]models.Post, len(args.PostIDs)),
		Custom:      pq.StringArray(args.Custom),
	}

	if pool.Custom == nil {
		pool.Custom = pq.StringArray([]string{})
	}

	for i, postID := range args.PostIDs {
		pool.Posts[i] = models.Post{
			ID: postID,
		}
	}

	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return pool, fmt.Errorf("sqlClient.BeginTx: %w", err)
	}
	defer tx.Rollback()

	err = r.poolQuery.Create(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("poolQuery.Create: %w", err)
	}

	err = r.poolPostQuery.AssociatePosts(ctx, tx, pool, pool.Posts)
	if err != nil {
		return pool, fmt.Errorf("poolPostQuery.AssociatePosts: %w", err)
	}

	err = r.postQuery.UpdatePoolCount(ctx, tx, pool.Posts, 1)
	if err != nil {
		return pool, fmt.Errorf("postQuery.UpdatePoolCount: %w", err)
	}

	err = r.poolQuery.GetFull(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("poolQuery.GetFull: %w", err)
	}

	tx.Commit()

	return pool, nil
}

func (r poolRepository) Delete(ctx context.Context, id int) (models.Pool, error) {
	pool := models.Pool{
		ID: id,
	}

	err := r.poolQuery.GetFull(ctx, r.sqlClient, &pool)
	if err != nil {
		return models.Pool{}, fmt.Errorf("poolRepository.GetFull: %w", err)
	}

	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return models.Pool{}, fmt.Errorf("sqlClient.BeginTx: %w", err)
	}
	defer tx.Rollback()

	err = r.poolQuery.Delete(
		ctx,
		tx,
		&models.Pool{
			ID: id,
		},
	)

	if err != nil {
		return models.Pool{}, fmt.Errorf("poolQuery.Delete: %w", err)
	}

	err = r.postQuery.UpdatePoolCount(ctx, tx, pool.Posts, -1)
	if err != nil {
		return models.Pool{}, fmt.Errorf("postQuery.UpdatePoolCount: %w", err)
	}

	tx.Commit()

	return pool, nil
}

func (r poolRepository) GetFull(ctx context.Context, id int) (models.Pool, error) {
	pool := models.Pool{
		ID: id,
	}

	err := r.poolQuery.GetFull(ctx, r.sqlClient, &pool)
	if err != nil {
		return models.Pool{}, fmt.Errorf("poolQuery.GetFull: %w", err)
	}

	return pool, nil
}

type PoolListFullArgs struct {
	Text     string
	Page     int
	PageSize int
}

func (r poolRepository) ListFull(ctx context.Context, args PoolListFullArgs) ([]models.Pool, int, error) {
	pools := make([]models.Pool, 0)
	count := 0

	err := r.poolQuery.ListFull(
		ctx,
		r.sqlClient,
		models.Search{
			Text:     args.Text,
			Page:     args.Page,
			PageSize: args.PageSize,
		},
		&pools,
		&count,
	)

	if err != nil {
		return pools, count, fmt.Errorf("poolQuery.ListFull: %w", err)
	}

	return pools, count, nil
}

type PoolUpdateArgs struct {
	Custom      *[]string
	Description *string
	ID          int
	Name        *string
	Posts       *[]int
}

func (r poolRepository) Update(ctx context.Context, args PoolUpdateArgs) (models.Pool, error) {
	pool := models.Pool{
		ID: args.ID,
	}

	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return pool, fmt.Errorf("sqlClient.BeginTx: %w", err)
	}
	defer tx.Rollback()

	err = r.poolQuery.GetFull(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("poolQuery.GetFull: %w", err)
	}

	if args.Custom != nil {
		pool.Custom = pq.StringArray(*args.Custom)
	}

	if args.Description != nil {
		pool.Description = *args.Description
	}

	if args.Name != nil {
		pool.Name = *args.Name
	}

	// TODO: don't allow to remove all posts
	if args.Posts != nil {
		// TODO: maybe should use int64 all over the place
		postIDs := make([]int64, len(*args.Posts))

		oldPostIDs := make([]int, len(pool.Posts))
		for i, post := range pool.Posts {
			oldPostIDs[i] = post.ID
		}

		toRemove := slice_utils.Difference(oldPostIDs, *args.Posts)
		toAdd := slice_utils.Difference(*args.Posts, oldPostIDs)

		pool.PostCount = len(*args.Posts)
		pool.Posts = make([]models.Post, pool.PostCount)

		for i, postID := range *args.Posts {
			pool.Posts[i] = models.Post{
				ID: postID,
			}
			postIDs[i] = int64(postID)
		}

		pool.PostIDs = postIDs

		if len(toRemove) > 0 {
			postsToRemove := make([]models.Post, len(toRemove))
			for i, postID := range toRemove {
				postsToRemove[i].ID = postID
			}

			err = r.poolPostQuery.DisassociatePosts(ctx, tx, pool, postsToRemove)
			if err != nil {
				return pool, fmt.Errorf("poolPostQuery.DisassociatePostsByID %w", err)
			}

			err = r.postQuery.UpdatePoolCount(ctx, tx, postsToRemove, -1)
			if err != nil {
				return pool, fmt.Errorf("postQuery.UpdatePoolCount: %w", err)
			}
		}

		if len(toAdd) > 0 {
			postsToAdd := make([]models.Post, len(toAdd))
			for i, postID := range toAdd {
				postsToAdd[i].ID = postID
			}

			err = r.poolPostQuery.AssociatePosts(ctx, tx, pool, postsToAdd)
			if err != nil {
				return pool, fmt.Errorf("poolPostQuery.AssociatePostsByID %w", err)
			}

			err = r.postQuery.UpdatePoolCount(ctx, tx, postsToAdd, 1)
			if err != nil {
				return pool, fmt.Errorf("postQuery.UpdatePoolCount: %w", err)
			}
		}

		//? always associate because the post position could have change
		err = r.poolPostQuery.AssociatePosts(ctx, tx, pool, pool.Posts)
		if err != nil {
			return pool, fmt.Errorf("poolPostQuery.AssociatePosts %w", err)
		}

		for i, postID := range *args.Posts {
			pool.Posts[i] = models.Post{
				ID: postID,
			}
		}
	}

	pool.UpdatedAt = time.Now()

	err = r.poolQuery.Update(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("poolQuery.Update: %w", err)
	}

	err = r.poolQuery.GetFull(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("poolQuery.GetFull: %w", err)
	}

	tx.Commit()

	return pool, nil
}
