package queries

import (
	"context"
	"fmt"
	"gobooru/internal/database"
	"gobooru/internal/models"

	"github.com/lib/pq"
)

type TagAliasQuery interface {
	// Will convert all tags in the tags slice to their alias if they have one.
	ResolveAlias(ctx context.Context, db database.DBClient, tags []string) error
}

type tagAliasQuery struct{}

func NewTagAliasQuery() TagAliasQuery {
	return &tagAliasQuery{}
}

func (q *tagAliasQuery) ResolveAlias(ctx context.Context, db database.DBClient, tags []string) error {
	var tagAliases []models.TagAlias

	err := db.SelectContext(
		ctx,
		&tagAliases,
		`
			SELECT tag_id, alias
			FROM tag_aliases
			WHERE alias = ANY($1)
		`,
		pq.Array(tags),
	)

	if err != nil {
		return fmt.Errorf("db.SelectContext: %w", err)
	}

	aliasMap := make(map[string]string, len(tagAliases))

	for _, tagAlias := range tagAliases {
		aliasMap[tagAlias.Alias] = tagAlias.TagID
	}

	for i, tag := range tags {
		if tagId, ok := aliasMap[tag]; ok {
			tags[i] = tagId
		}
	}

	return nil
}
