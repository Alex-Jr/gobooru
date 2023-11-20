package queries

import (
	"context"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"time"
)

type PostQuery interface {
	Create(ctx context.Context, db database.DBClient, post *models.Post) error
	Delete(ctx context.Context, db database.DBClient, post *models.Post) error
	GetFull(ctx context.Context, db database.DBClient, post *models.Post) error
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
				"description",
				"updated_at"
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
		return err
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
		return err
	}

	return nil
}
