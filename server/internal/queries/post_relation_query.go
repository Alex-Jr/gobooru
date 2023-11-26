package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"
)

type PostRelationQuery interface {
	InsertRelations(ctx context.Context, db database.DBClient, post models.Post, relations []models.PostRelation) error
}

type postRelationQuery struct{}

func NewPostRelationQuery() PostRelationQuery {
	return &postRelationQuery{}
}

func (q *postRelationQuery) InsertRelations(ctx context.Context, db database.DBClient, post models.Post, relations []models.PostRelation) error {
	if len(relations) == 0 {
		return nil
	}

	_, err := db.NamedExecContext(
		ctx,
		`
			INSERT INTO "post_relations" (
				"created_at",
				"post_id", 
				"other_post_id",
				"similarity",
				"type"
			) VALUES (
				:created_at,
				:post_id,
				:other_post_id,
				:similarity,
				:type
			) ON CONFLICT DO NOTHING`,
		relations,
	)

	if err != nil {
		return fmt.Errorf("db.NamedExecContext: %w", err)
	}

	return nil
}
