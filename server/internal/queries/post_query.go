package queries

import (
	"context"
	"database/sql"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/query_parser"
	"time"

	"github.com/lib/pq"
)

type PostQuery interface {
	Create(ctx context.Context, db database.DBClient, post *models.Post) error
	Delete(ctx context.Context, db database.DBClient, post *models.Post) error
	GetFull(ctx context.Context, db database.DBClient, post *models.Post) error
	List(ctx context.Context, db database.DBClient, search models.Search, posts *[]models.Post, count *int) error
	Update(ctx context.Context, db database.DBClient, post models.Post) error
	UpdatePoolCount(ctx context.Context, db database.DBClient, post []models.Post, increment int) error
	RemoveTag(ctx context.Context, db database.DBClient, tag string) error
}

type postQuery struct {
	queryParser query_parser.Parser
}

func NewPostQuery() PostQuery {
	return &postQuery{
		queryParser: query_parser.NewParser(query_parser.ParserConfig{
			WhereField: map[string]query_parser.WhereField{
				"tag": {
					DBName:   "pt.\"tag_ids\"",
					Operator: "@>",
					ParserFn: query_parser.ArrayParserFn,
				},
			},
			SortField: map[string]query_parser.SortField{
				"id": {
					DBName:       "pt.\"id\"",
					DefaultOrder: "DESC",
				},
			},
			DefaultWhereField: "tag",
			DefaultSortField:  "id",
		}),
	}
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
				"md5",
				"file_ext",
				"file_size",
				"file_path",
				"thumb_path",
				"created_at",
				"updated_at"
			) VALUES (
				:rating,
				:description,
				:tag_ids,
				:tag_count,
				:pool_count,
				:md5,
				:file_ext,
				:file_size,
				:file_path,
				:thumb_path,
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
			WITH 
				"pools" AS (
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
				),
				"tags" AS (
					SELECT
						$1::int as "post_id",
						JSONB_AGG(to_jsonb(t.*)) as "tags"
					FROM
						"tags" "t"
					INNER JOIN "post_tags" "pt" ON
						"pt"."tag_id" = "t"."id"
					WHERE
						"pt"."post_id" = $1::int
					GROUP BY
						"pt"."post_id"
				),
				"relations" AS (
					SELECT
						$1::int as "post_id",
						JSONB_AGG(
							TO_JSONB(pr.*) || JSONB_BUILD_OBJECT('other_post', TO_JSONB(op.*))
						) AS "relations"
					FROM
						"posts" op
					INNER JOIN "post_relations" "pr" ON
						"pr"."other_post_id" = "op"."id"
					WHERE
						"pr"."post_id" = $1::int
				)
			SELECT
				p."created_at",
				p."description",
				p."id",
				p."updated_at",
				p."rating",
				p."tag_count",
				p."tag_ids",
				p."pool_count",
				p."md5",
				p."file_ext",
				p."file_size",
				p."file_path",
				p."thumb_path",
				pl."pools",
				t."tags",
				r."relations"
			FROM
				"posts" as "p"
			LEFT JOIN "pools" as "pl" ON
				"pl"."post_id" = "p"."id"
			LEFT JOIN "tags" as "t" ON
				"t"."post_id" = "p"."id"
			LEFT JOIN "relations" AS "r" ON
				"r"."post_id" = "p"."id"
			WHERE
				p."id" = $1::int
		`,
		post.ID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return database.ErrNotFound
		}
		return fmt.Errorf("db.GetContext: %w", err)
	}

	return nil
}

func (q *postQuery) List(ctx context.Context, db database.DBClient, search models.Search, posts *[]models.Post, count *int) error {
	parsed, err := q.queryParser.ParseSearch(search)
	if err != nil {
		return fmt.Errorf("queryParser.ParseSearch: %w", err)
	}

	err = db.GetContext(
		ctx,
		count,
		`
			SELECT
				count(*)
			FROM
				"posts" pt
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
		posts,
		fmt.Sprintf(
			`
				SELECT
					pt."created_at",
					pt."description",
					pt."id",
					pt."pool_count",
					pt."rating",
					pt."tag_count",
					pt."tag_ids",
					pt."md5",
					pt."file_ext",
					pt."file_size",
					pt."file_path",
					pt."thumb_path",
					pt."updated_at"
				FROM
					"posts" pt
				WHERE
					%s
				GROUP BY
					pt."id"
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

func (q *postQuery) Update(ctx context.Context, db database.DBClient, post models.Post) error {
	post.UpdatedAt = time.Now()

	_, err := db.NamedExecContext(
		ctx,
		`
			UPDATE
				"posts"
			SET
				"rating" = :rating,
				"description" = :description,
				"tag_ids" = :tag_ids,
				"tag_count" = :tag_count,
				"pool_count" = :pool_count,
				"updated_at" = :updated_at
			WHERE
				"id" = :id
		`,
		post,
	)

	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}

	return nil
}

func (q *postQuery) RemoveTag(ctx context.Context, db database.DBClient, tag string) error {
	_, err := db.ExecContext(
		ctx,
		`
			UPDATE
				"posts"
			SET
				"tag_count" = "tag_count" - 1,
				"tag_ids" = array_remove("tag_ids", $1)
			WHERE
				"tag_ids" @> ARRAY[$1]
		`,
		tag,
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
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
