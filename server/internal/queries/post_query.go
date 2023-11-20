package queries

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"time"
)

type PostQuery interface {
	Create(ctx context.Context, db database.DBClient, post *models.Post) error
}

type postQuery struct {
}

func NewPostQuery() PostQuery {
	return &postQuery{}
}

func (q *postQuery) Create(ctx context.Context, db database.DBClient, post *models.Post) error {
	now := time.Now()

	post.CreatedAt = now
	post.UpdatedAt = now

	rows, err := db.NamedQueryContext(
		ctx,
		`
			INSERT INTO "posts" (
				"created_at",
				"description"
				"updated_at",
			) VALUES (
				:created_at,
				:description,
				:updated_at
			) RETURNING 
				"id"
		`,
		post,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(post)

		if err != nil {
			return err
		}
	}

	return nil
}
