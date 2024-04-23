package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
	"gobooru/internal/query_parser"

	"github.com/lib/pq"
)

type TagQuery interface {
	CreateMany(ctx context.Context, db database.DBClient, tags *[]models.Tag) error
	Delete(ctx context.Context, db database.DBClient, tag models.Tag) error
	Get(ctx context.Context, db database.DBClient, tag *models.Tag) error
	List(ctx context.Context, db database.DBClient, search models.Search, count *int, tags *[]models.Tag) error
	UpdatePostCount(ctx context.Context, db database.DBClient, tags []string, increment int) error
}

type tagQuery struct {
	parser query_parser.Parser
}

func NewTagQuery() TagQuery {
	return &tagQuery{
		parser: query_parser.NewParser(query_parser.ParserConfig{
			WhereField: map[string]query_parser.WhereField{
				"id": {
					DBName:   "t.\"id\"",
					Operator: "ILIKE",
					ParserFn: query_parser.LikeParserFnConfigurable(query_parser.RIGHT),
				},
				"post_count": {
					DBName:   "t.\"post_count\"",
					Rangable: true,
					ParserFn: query_parser.IntParserFn,
				},
			},
			SortField: map[string]query_parser.SortField{
				"id": {
					DBName:       "t.\"id\"",
					DefaultOrder: "ASC",
				},
				"post_count": {
					DBName:       "t.\"post_count\"",
					DefaultOrder: "DESC",
				},
			},
			DefaultWhereField: "id",
			DefaultSortField:  "id",
		}),
	}
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

func (q *tagQuery) List(ctx context.Context, db database.DBClient, search models.Search, count *int, tags *[]models.Tag) error {
	parsed, err := q.parser.ParseSearch(search)
	if err != nil {
		return fmt.Errorf("parser.ParseSearch: %w", err)
	}

	err = db.GetContext(
		ctx,
		count,
		fmt.Sprintf(
			`
				SELECT
					COUNT(*)
				FROM
					"tags" t
				WHERE
					%s
			`,
			parsed.WhereQuery,
		),
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
		tags,
		fmt.Sprintf(`
			SELECT
				t."id",
				t."description",
				t."post_count",
				t."category_id",
				t."created_at",
				t."updated_at"
			FROM
				"tags" t
			WHERE
				%s
			ORDER BY
				%s`,
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
