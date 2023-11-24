package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"time"

	"github.com/lib/pq"
)

type PoolPostQuery interface {
	// AssociatePosts associates posts to a pool.
	AssociatePosts(ctx context.Context, db database.DBClient, pool models.Pool, posts []models.Post) error
	// DisassociatePosts disassociates posts from a pool.
	DisassociatePosts(ctx context.Context, db database.DBClient, pool models.Pool, posts []models.Post) error
}

type poolPostQuery struct {
}

func NewPoolPostQuery() PoolPostQuery {
	return &poolPostQuery{}
}

func (p poolPostQuery) AssociatePosts(ctx context.Context, db database.DBClient, pool models.Pool, posts []models.Post) error {
	now := time.Now()

	poolPosts := make([]models.PoolPost, len(posts))
	for i, posts := range posts {
		poolPosts[i] = models.PoolPost{
			PoolID:    pool.ID,
			PostID:    posts.ID,
			Position:  i,
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	_, err := db.NamedExecContext(
		ctx,
		`
			INSERT INTO pool_posts (pool_id, post_id, position, created_at, updated_at)
			VALUES (:pool_id, :post_id, :position, :created_at, :updated_at)
			ON CONFLICT (pool_id, post_id) DO UPDATE SET
				position = EXCLUDED.position,
				updated_at = EXCLUDED.updated_at
		`,
		poolPosts,
	)

	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}

	return nil
}

func (p poolPostQuery) DisassociatePosts(ctx context.Context, db database.DBClient, pool models.Pool, posts []models.Post) error {
	postIDs := make([]int, len(posts))

	for i, post := range posts {
		postIDs[i] = post.ID
	}

	_, err := db.ExecContext(
		ctx,
		`
			DELETE FROM
				"pool_posts"
			WHERE
				"pool_id" = $1
				AND "post_id" = ANY($2)
		`,
		pool.ID,
		pq.Array(postIDs),
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}
