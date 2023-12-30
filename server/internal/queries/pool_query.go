package queries

import (
	"context"
	"database/sql"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/query_parser"
	"time"
)

type PoolQuery interface {
	// Create creates a pool entry.
	Create(ctx context.Context, db database.DBClient, pool *models.Pool) error
	// Delete	deletes a pool entry.
	Delete(ctx context.Context, db database.DBClient, pool *models.Pool) error
	// GetFull fetches a pool by ID and all of its posts.
	GetFull(ctx context.Context, db database.DBClient, pool *models.Pool) error
	RemovePost(ctx context.Context, db database.DBClient, postID int) error
	// ListFull fetches all pools and their posts based on the given arguments.
	ListFull(ctx context.Context, db database.DBClient, search models.Search, pools *[]models.Pool, count *int) error
	// Update updates a pool entry.
	Update(ctx context.Context, db database.DBClient, pool *models.Pool) error
}

type poolQuery struct {
	parser query_parser.Parser
}

func NewPoolQuery() PoolQuery {
	return &poolQuery{
		parser: query_parser.NewParser(query_parser.ParserConfig{
			WhereField: map[string]query_parser.WhereField{
				"id": {
					DBName:   "pl.id",
					Operator: "=",
					ParserFn: query_parser.IntParserFn,
				},
				"name": {
					DBName:   "pl.name",
					Operator: "ILIKE",
					ParserFn: query_parser.LikeParserFn,
				},
				"custom": {
					DBName:   "pl.custom",
					Operator: "@>",
					ParserFn: query_parser.ArrayParserFn,
				},
				"createdAt": {
					DBName:   "pl.created_at",
					Rangable: true,
					ParserFn: query_parser.TimeParserFn,
				},
				"postCount": {
					DBName:   "pl.post_count",
					Rangable: true,
					ParserFn: query_parser.IntParserFn,
				},
			},
			SortField: map[string]query_parser.SortField{
				"id": {
					DBName:       "pl.id",
					DefaultOrder: "DESC",
				},
			},
			DefaultWhereField: "name",
			DefaultSortField:  "id",
		}),
	}
}

func (q poolQuery) Create(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	now := time.Now()

	pool.CreatedAt = now
	pool.UpdatedAt = now

	rows, err := db.NamedQueryContext(
		ctx,
		`
			INSERT INTO pools (
				"created_at",
				"custom",
				"description", 
				"name", 
				"post_ids",
				"post_count",
				"updated_at"
			)
			VALUES (
				:created_at,
				:custom,
				:description,
				:name, 
				:post_ids,
				:post_count,
				:updated_at
			)
			RETURNING 
				"id"
		`,
		pool,
	)
	if err != nil {
		return fmt.Errorf("db.NamedQueryContext: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(pool)
		if err != nil {
			return fmt.Errorf("rows.StructScan: %w", err)
		}
	}

	return nil
}

func (q poolQuery) Delete(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	_, err := db.ExecContext(
		ctx,
		`
			DELETE FROM
				"pools"
			WHERE
				"id" = $1
		`,
		pool.ID,
	)

	if err != nil {
		return fmt.Errorf(" db.ExecContext: %w", err)
	}

	return nil
}

func (q poolQuery) GetFull(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	err := db.GetContext(
		ctx,
		pool,
		`
			SELECT 
				pl."created_at", 
				pl."custom", 
				pl."description", 
				pl."id",
				pl."name", 
				JSONB_AGG(
					ROW_TO_JSON(pt.*)
					ORDER BY pp."position"
				) AS "posts",
				pl."post_count",
				pl."post_ids",
				pl."updated_at"
			FROM
				"pools" pl
			INNER JOIN "pool_posts" pp ON
				pp."pool_id" = pl."id"
			INNER JOIN "posts" pt ON
				pp."post_id" = pt."id" 
			WHERE
				pl."id" = $1
			GROUP BY
				pl."id"
		`,
		pool.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return database.ErrNotFound
		}

		return fmt.Errorf("db.GetContext: %w", err)
	}

	return nil
}

func (q poolQuery) RemovePost(ctx context.Context, db database.DBClient, postID int) error {
	_, err := db.ExecContext(
		ctx,
		`
		UPDATE "pools"
		SET
			"post_count" = "post_count" - 1,
			"post_ids" = array_remove("post_ids", $1)
		WHERE
			$1 = ANY("post_ids")
		`,
		postID,
	)
	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}

func (q poolQuery) ListFull(ctx context.Context, db database.DBClient, search models.Search, pools *[]models.Pool, count *int) error {
	parsed, err := q.parser.ParseSearch(search)
	if err != nil {
		return fmt.Errorf("parser.ParseSearch: %w", err)
	}

	err = db.GetContext(
		ctx,
		count,
		`
			SELECT
				count(*)
			FROM "pools" pl
			WHERE
		`+parsed.WhereQuery,
		parsed.WhereArgs...,
	)
	if err != nil {
		return fmt.Errorf("db.GetContext: %w", err)
	}

	// TODO: maybe make count and list parallel
	if count == nil || *count == 0 {
		return nil
	}

	err = db.SelectContext(
		ctx,
		pools,
		fmt.Sprintf(`
			SELECT 
				pl."created_at", 
				pl."custom", 
				pl."description", 
				pl."id",
				pl."name", 
				pl."post_count",
				JSONB_AGG(
					ROW_TO_JSON(pt.*)
					ORDER BY pp."position"
				) AS "posts",
				pl."post_ids",
				pl."updated_at"
			FROM
				"pools" pl
			INNER JOIN "pool_posts" pp ON
				pp."pool_id" = pl."id"
			INNER JOIN "posts" pt ON
				pp."post_id" = pt."id" 
			WHERE
				%s
			GROUP BY
				pl."id"
			ORDER BY
				%s
			`,
			parsed.WhereQuery,
			parsed.SortQuery,
		),
		append(parsed.WhereArgs, parsed.PaginationArgs...)...,
	)
	if err != nil {
		return fmt.Errorf("db.SelectContext: %w", err)
	}

	return nil
}

func (q poolQuery) Update(ctx context.Context, db database.DBClient, pool *models.Pool) error {
	_, err := db.NamedExecContext(
		ctx,
		`
			UPDATE
				"pools"
			SET
				"custom" = :custom,
				"description" = :description,
				"name" = :name,
				"post_count" = :post_count,
				"updated_at" = :updated_at
			WHERE
				"id" = :id
		`,
		pool,
	)
	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}

	return nil
}
