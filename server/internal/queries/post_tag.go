package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"time"

	"github.com/lib/pq"
)

type PostTagQuery interface {
	// AssociatePosts associates tags to a post.
	AssociatePosts(ctx context.Context, db database.DBClient, post models.Post, tags []models.Tag) error
	// DisassociatePosts disassociates posts from a pool.
	DisassociatePostsByID(ctx context.Context, db database.DBClient, post models.Post, tagIDs []string) error
}

type postTagQuery struct {
}

func NewPostTagQuery() PostTagQuery {
	return &postTagQuery{}
}

func (q postTagQuery) AssociatePosts(ctx context.Context, db database.DBClient, post models.Post, tags []models.Tag) error {
	now := time.Now()

	postTags := make([]models.PostTag, len(tags))

	for i, tag := range tags {
		postTags[i] = models.PostTag{
			PostID:    post.ID,
			TagID:     tag.ID,
			CreatedAt: now,
			UpdatedAt: now,
		}
	}

	_, err := db.NamedExecContext(
		ctx,
		`
			INSERT INTO post_tags (post_id, tag_id, created_at, updated_at)
			VALUES (:post_id, :tag_id, :created_at, :updated_at)
			ON CONFLICT (post_id, tag_id) DO NOTHING
		`,
		postTags,
	)

	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}

	return nil

}

func (q postTagQuery) DisassociatePostsByID(ctx context.Context, db database.DBClient, post models.Post, tagIDs []string) error {
	_, err := db.ExecContext(
		ctx,
		`
			DELETE FROM "post_tags"
			WHERE 
				"post_id" = $1
				AND "tag_id" = ANY($2)
		`,
		post.ID,
		pq.Array(tagIDs),
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}
