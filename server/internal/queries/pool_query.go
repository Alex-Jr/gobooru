package queries

import (
	"context"
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
			},
			SortField: map[string]query_parser.SortField{
				"id": {
					DBName:       "pl.id",
					DefaultOrder: "DESC",
				},
			},
			DefaultWhereField: "id",
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
				"description", 
				"name", 
				"post_count",
				"updated_at"
			)
			VALUES (
				:created_at,
				:description,
				:name, 
				:post_count,
				:updated_at
			)
			RETURNING 
				"id"
		`,
		pool,
	)
	if err != nil {
		return fmt.Errorf("creating pool: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(pool)
		if err != nil {
			return fmt.Errorf("scanning pool: %w", err)
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
		return fmt.Errorf("deleting pool: %w", err)
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
		return fmt.Errorf("finding pool: %w", err)
	}

	return nil
}

func (q poolQuery) ListFull(ctx context.Context, db database.DBClient, search models.Search, pools *[]models.Pool, count *int) error {
	parsed, err := q.parser.ParseSearch(search)
	if err != nil {
		return fmt.Errorf("parsing search: %w", err)
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
		return fmt.Errorf("counting pools: %w", err)
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
				JSONB_AGG(
					ROW_TO_JSON(pt.*)
					ORDER BY pp."position"
				) AS "posts",
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
		return fmt.Errorf("listing pools: %w", err)
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
		return fmt.Errorf("updating pool: %w", err)
	}

	return nil
}
