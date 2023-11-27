package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"

	"github.com/lib/pq"
)

type TagQuery interface {
	CreateMany(ctx context.Context, db database.DBClient, tags *[]models.Tag) error
	Delete(ctx context.Context, db database.DBClient, tag models.Tag) error
	Get(ctx context.Context, db database.DBClient, tag *models.Tag) error
	UpdatePostCount(ctx context.Context, db database.DBClient, tags []string, increment int) error
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
				"category_id",
				"created_at", 
				"updated_at"
			) VALUES (
				:id,
				:description,
				:post_count,
				:category_id,
				:created_at,
				:updated_at
			)
			ON CONFLICT ("id") DO UPDATE SET
				post_count = "tags".post_count + 1
			RETURNING
				"post_count",
				"created_at",
				"updated_at"
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

func (q *tagQuery) Delete(ctx context.Context, db database.DBClient, tag models.Tag) error {
	_, err := db.ExecContext(
		ctx,
		`
			DELETE FROM
				"tags"
			WHERE
				id = $1
		`,
		tag.ID,
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}

func (q *tagQuery) Get(ctx context.Context, db database.DBClient, tag *models.Tag) error {
	err := db.GetContext(
		ctx,
		tag,
		`
			SELECT
				"id",
				"description",
				"post_count",
				"category_id",
				"created_at",
				"updated_at"
			FROM
				"tags"
			WHERE
				id = $1
		`,
		tag.ID,
	)

	if err != nil {
		return fmt.Errorf("db.GetContext: %w", err)
	}

	return nil
}

func (q *tagQuery) UpdatePostCount(ctx context.Context, db database.DBClient, tags []string, increment int) error {
	if len(tags) == 0 {
		return nil
	}

	_, err := db.ExecContext(
		ctx,
		`
			UPDATE "tags"
			SET
				post_count = post_count + $2
			WHERE
				id = ANY($1)
			RETURNING
				"post_count"
		`,
		pq.Array(tags),
		increment,
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}
