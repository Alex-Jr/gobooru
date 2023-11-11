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
	AssociatePosts(ctx context.Context, db database.DBClient, poolID int, posts []models.Post) error
	// DisassociatePosts disassociates posts from a pool.
	DisassociatePostsByID(ctx context.Context, db database.DBClient, poolID int, postIDs []int) error
}

type poolPostQuery struct {
}

func NewPoolPostQuery() PoolPostQuery {
	return &poolPostQuery{}
}

func (p poolPostQuery) AssociatePosts(ctx context.Context, db database.DBClient, poolID int, posts []models.Post) error {
	now := time.Now()

	poolPosts := make([]models.PoolPost, len(posts))
	for i, posts := range posts {
		poolPosts[i] = models.PoolPost{
			PoolID:    poolID,
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
		return fmt.Errorf("creating pool posts: %w", err)
	}

	return nil
}

func (p poolPostQuery) DisassociatePostsByID(ctx context.Context, db database.DBClient, poolID int, postIDs []int) error {
	_, err := db.ExecContext(
		ctx,
		`
			DELETE FROM
				"pool_posts"
			WHERE
				"pool_id" = $1
				AND "post_id" = ANY($2)
		`,
		poolID,
		pq.Array(postIDs),
	)

	if err != nil {
		return fmt.Errorf("deleting pool posts: %w", err)
	}

	return nil
}
