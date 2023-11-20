package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
)

type TagQuery interface {
	CreateMany(ctx context.Context, db database.DBClient, tags *[]models.Tag) error
}

type tagQuery struct {
}

func NewTagQuery() TagQuery {
	return &tagQuery{}
}

func (q *tagQuery) CreateMany(ctx context.Context, db database.DBClient, tags *[]models.Tag) error {
	if len(*tags) == 0 {
		return nil
	}

	rows, err := db.NamedQueryContext(
		ctx,
		`
			INSERT INTO "tags" (
				"id", 
				"description", 
				"post_count", 
				"created_at", 
				"updated_at"
			) VALUES (
				:id,
				:description,
				:post_count,
				:created_at,
				:updated_at
			)
			ON CONFLICT ("id") DO UPDATE SET
				post_count = "tags".post_count + 1
			RETURNING
				"post_count"
		`,
		*tags,
	)

	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		err = rows.StructScan(&(*tags)[i])

		if err != nil {
			return fmt.Errorf("rows.StructScan: %w", err)
		}

		i++
	}

	return nil
}
