package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"time"
)

type PostNotesQuery interface {
	// Create creates a post note entry.
	Create(ctx context.Context, db database.DBClient, postNote *models.PostNote) error
}

type postNotesQuery struct {
}

func NewPostNotesQuery() PostNotesQuery {
	return &postNotesQuery{}
}

func (q *postNotesQuery) Create(ctx context.Context, db database.DBClient, postNote *models.PostNote) error {
	now := time.Now()

	postNote.CreatedAt = now
	postNote.UpdatedAt = now

	_, err := db.NamedExecContext(
		ctx,
		`
			INSERT INTO post_notes (post_id, body, x, y, width, height, created_at, updated_at)
			VALUES (:post_id, :body, :x, :y, :width, :height, :created_at, :updated_at)
			RETURNING "id"
		`,
		postNote,
	)
	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}

	return nil
}
