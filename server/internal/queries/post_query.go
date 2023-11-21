package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"time"

	"github.com/lib/pq"
)

type PostQuery interface {
	Create(ctx context.Context, db database.DBClient, post *models.Post) error
	Delete(ctx context.Context, db database.DBClient, post *models.Post) error
	GetFull(ctx context.Context, db database.DBClient, post *models.Post) error
	UpdatePoolCount(ctx context.Context, db database.DBClient, post []models.Post, increment int) error
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
				"rating",
				"description",
				"tag_ids",
				"tag_count",
				"pool_count",
				"created_at",
				"updated_at"
			) VALUES (
				:rating,
				:description,
				:tag_ids,
				:tag_count,
				:pool_count,
				:created_at,
				:updated_at
			) RETURNING 
				"id"
		`,
		post,
	)

	if err != nil {
		return fmt.Errorf("db.NamedQueryContext: %w", err)
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(post)

		if err != nil {
			return fmt.Errorf("rows.StructScan: %w", err)
		}
	}

	return nil
}

func (q *postQuery) Delete(ctx context.Context, db database.DBClient, post *models.Post) error {
	_, err := db.ExecContext(
		ctx,
		`
			DELETE FROM
				"posts"
			WHERE
				"id" = $1
		`,
		post.ID,
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}

func (q *postQuery) GetFull(ctx context.Context, db database.DBClient, post *models.Post) error {
	err := db.GetContext(
		ctx,
		post,
		`
			WITH "pools" AS (
				SELECT
					$1::int as "post_id",
					JSONB_AGG(to_jsonb(pl.*)) as "pools"
				FROM
					"pools" "pl"
				INNER JOIN "pool_posts" "pp" ON
					"pp"."pool_id" = "pl"."id"
				WHERE
					"pp"."post_id" = $1::int
				GROUP BY
					"pp"."post_id"
			)
			SELECT
				p."created_at",
				p."description",
				p."id",
				p."updated_at",
				pl."pools"
			FROM
				"posts" as "p"
			LEFT JOIN "pools" as "pl" ON
				"pl"."post_id" = "p"."id"
			WHERE
				p."id" = $1::int
		`,
		post.ID,
	)

	if err != nil {
		return fmt.Errorf("db.GetContext: %w", err)
	}

	return nil
}

func (q *postQuery) UpdatePoolCount(ctx context.Context, db database.DBClient, posts []models.Post, increment int) error {
	postIDs := make([]int, len(posts))

	for i, post := range posts {
		postIDs[i] = post.ID
	}

	_, err := db.ExecContext(
		ctx,
		`
			UPDATE
				"posts"
			SET
				"pool_count" = "pool_count" + $1
			WHERE
				"id" = ANY($2)
		`,
		increment,
		pq.Array(postIDs),
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}
