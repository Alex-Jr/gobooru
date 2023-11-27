package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
)

type TagCategoryQuery interface {
	UpdateTagCount(ctx context.Context, db database.DBClient, tagCategory string, value int) error
}

type tagCategoryQuery struct {
}

func NewTagCategoryQuery() TagCategoryQuery {
	return &tagCategoryQuery{}
}

func (q *tagCategoryQuery) UpdateTagCount(ctx context.Context, db database.DBClient, tagCategory string, value int) error {
	_, err := db.ExecContext(
		ctx,
		`
			UPDATE "tag_categories"
			SET "tag_count" = "tag_categories"."tag_count" + $1
			WHERE "tag_categories"."id" = $2
		`,
		value,
		tagCategory,
	)

	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}
