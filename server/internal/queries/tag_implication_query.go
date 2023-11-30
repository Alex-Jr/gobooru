package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"

	"github.com/lib/pq"
)

type TagImplicationQuery interface {
	// Will add all implied tags to the tags slice.
	ResolveImplications(ctx context.Context, db database.DBClient, tags *[]string) error
}

type tagImplicationQuery struct{}

func NewTagImplicationQuery() TagImplicationQuery {
	return &tagImplicationQuery{}
}

func (q *tagImplicationQuery) ResolveImplications(ctx context.Context, db database.DBClient, tags *[]string) error {
	var tagImplications []models.TagImplication

	err := db.SelectContext(
		ctx,
		&tagImplications,
		`
			SELECT 
				"implication_id"
			FROM 
				"tag_implications"
			WHERE 
				"tag_id" = ANY($1)
		`,
		pq.Array(tags),
	)
	if err != nil {
		return fmt.Errorf("db.SelectContext: %w", err)
	}

	// using map to deduplicate the implications
	tagsMap := make(map[string]bool)

	for i := range tagImplications {
		tagsMap[tagImplications[i].ImplicationID] = true
	}

	for i := range tagsMap {
		*tags = append(*tags, i)
	}

	return nil
}
