package repositories

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/queries"
	"gobooru/internal/slice_utils"

	"github.com/lib/pq"
)

type PoolRepository interface {
	Create(ctx context.Context, args PoolCreateArgs) (models.Pool, error)
	Delete(ctx context.Context, id int) error
	GetFull(ctx context.Context, id int) (models.Pool, error)
	ListFull(ctx context.Context, args PoolListFullArgs) ([]models.Pool, int, error)
	Update(ctx context.Context, args PoolUpdateArgs) (models.Pool, error)
}

type poolRepository struct {
	sqlClient     database.SQLClient
	poolQuery     queries.PoolQuery
	poolPostQuery queries.PoolPostQuery
}

func NewPoolRepository(dbClient database.SQLClient) PoolRepository {
	return &poolRepository{
		sqlClient:     dbClient,
		poolQuery:     queries.NewPoolQuery(),
		poolPostQuery: queries.NewPoolPostQuery(),
	}
}

type PoolCreateArgs struct {
	Custom      []string
	Description string
	Name        string
	Posts       []int
}

func (r poolRepository) Create(ctx context.Context, args PoolCreateArgs) (models.Pool, error) {

	pool := models.Pool{
		Description: args.Description,
		ID:          0,
		Name:        args.Name,
		PostCount:   len(args.Posts),
		Posts:       make([]models.Post, len(args.Posts)),
		Custom:      pq.StringArray(args.Custom),
	}

	if pool.Custom == nil {
		pool.Custom = pq.StringArray([]string{})
	}

	for i, postID := range args.Posts {
		pool.Posts[i] = models.Post{
			ID: postID,
		}
	}

	tx, err := r.sqlClient.BeginTxx(ctx, nil)
	if err != nil {
		return pool, fmt.Errorf("beginning transaction: %w", err)
	}
	defer tx.Rollback()

	err = r.poolQuery.Create(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("creating pool: %w", err)
	}

	err = r.poolPostQuery.AssociatePosts(ctx, tx, pool.ID, pool.Posts)
	if err != nil {
		return pool, fmt.Errorf("creating pool posts: %w", err)
	}

	err = r.poolQuery.GetFull(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("finding posts: %w", err)
	}

	tx.Commit()

	return pool, nil
}

func (r poolRepository) Delete(ctx context.Context, id int) error {
	err := r.poolQuery.Delete(ctx, r.sqlClient, &models.Pool{
		ID: id,
	})

	if err != nil {
		return fmt.Errorf("deleting pool: %w", err)
	}

	return nil
}

func (r poolRepository) GetFull(ctx context.Context, id int) (models.Pool, error) {
	pool := models.Pool{
		ID: id,
	}

	err := r.poolQuery.GetFull(ctx, r.sqlClient, &pool)
	if err != nil {
		return pool, fmt.Errorf("finding pool: %w", err)
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
		return pools, count, fmt.Errorf("listing pools: %w", err)
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
		return pool, fmt.Errorf("beginning transaction: %w", err)
	}
	defer tx.Rollback()

	err = r.poolQuery.GetFull(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("finding pool: %w", err)
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

	if args.Posts != nil {
		oldPostIDs := make([]int, len(pool.Posts))
		for i, post := range pool.Posts {
			oldPostIDs[i] = post.ID
		}

		toRemove := slice_utils.Difference(*args.Posts, oldPostIDs)

		pool.PostCount = len(*args.Posts)
		pool.Posts = make([]models.Post, pool.PostCount)

		for i, postID := range *args.Posts {
			pool.Posts[i] = models.Post{
				ID: postID,
			}
		}

		if len(toRemove) > 0 {
			err = r.poolPostQuery.DisassociatePostsByID(ctx, tx, pool.ID, toRemove)
			if err != nil {
				return pool, fmt.Errorf("deleting pool posts: %w", err)
			}
		}

		//? always associate because the post position could have change
		err = r.poolPostQuery.AssociatePosts(ctx, tx, pool.ID, pool.Posts)
		if err != nil {
			return pool, fmt.Errorf("creating pool posts: %w", err)
		}

		for i, postID := range *args.Posts {
			pool.Posts[i] = models.Post{
				ID: postID,
			}
		}
	}

	err = r.poolQuery.Update(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("updating pool: %w", err)
	}

	err = r.poolQuery.GetFull(ctx, tx, &pool)
	if err != nil {
		return pool, fmt.Errorf("finding posts: %w", err)
	}

	tx.Commit()

	return pool, nil
}
