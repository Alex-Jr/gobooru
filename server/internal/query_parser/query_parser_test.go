package query_parser_test

import (
	"testing"

	"gobooru/internal/models"
	"gobooru/internal/query_parser"

	"github.com/stretchr/testify/assert"
)

func TestParser_SearchParser(t *testing.T) {
	parser := query_parser.NewParser(query_parser.ParserConfig{
		WhereField: map[string]query_parser.WhereField{
			"tag": {
				DBName:   "posts.tags",
				Operator: "@>",
			},
			"rating": {
				DBName:   "posts.rating",
				Operator: "=",
			},
			"score": {
				DBName:   "posts.score",
				Rangable: true,
			},
			"createdAt": {
				DBName:   "posts.created_at",
				Rangable: true,
				ParserFn: query_parser.TimeParserFn,
			},
		},
		SortField: map[string]query_parser.SortField{
			"score": {
				Name:         "score",
				DBName:       "posts.score",
				AllowASC:     true,
				AllowDESC:    true,
				DefaultOrder: "DESC",
			},
			"fav": {
				Name:         "fav",
				DBName:       "posts.fav_count",
				DefaultOrder: "DESC",
			},
		},
		DefaultWhereField: "tag",
		DefaultSortField:  "score",
	})

	tests := []struct {
		name   string
		search models.Search
		want   []interface{}
	}{
		{
			name: "empty search",
			search: models.Search{
				Text:     "",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"1 = 1",
				"posts.score DESC LIMIT $1 OFFSET $2",
				[]interface{}{10, 0},
			},
		},
		{
			name: "one word",
			search: models.Search{
				Text:     "tag:cat",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.tags @> $1 )",
				"posts.score DESC LIMIT $2 OFFSET $3",
				[]interface{}{"cat", 10, 0},
			},
		},
		{
			name: "multiple words",
			search: models.Search{
				Text:     "tag:cat rating:safe",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.tags @> $1 ) AND ( posts.rating = $2 )",
				"posts.score DESC LIMIT $3 OFFSET $4",
				[]interface{}{"cat", "safe", 10, 0},
			},
		},
		{
			name: "multiple words with comma",
			search: models.Search{
				Text:     "tag:cat rating:safe,questionable",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.tags @> $1 ) AND ( posts.rating = $2 OR posts.rating = $3 )",
				"posts.score DESC LIMIT $4 OFFSET $5",
				[]interface{}{"cat", "safe", "questionable", 10, 0},
			},
		},
		{
			name: "non existing field",
			search: models.Search{
				Text:     "non-existing:cat",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"1 = 1",
				"posts.score DESC LIMIT $1 OFFSET $2",
				[]interface{}{10, 0},
			},
		},
		{
			name: "rangable field",
			search: models.Search{
				Text:     "score:1..10",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.score BETWEEN $1 AND $2 )",
				"posts.score DESC LIMIT $3 OFFSET $4",
				[]interface{}{"1", "10", 10, 0},
			},
		},
		{
			name: "rangable field greater than",
			search: models.Search{
				Text:     "score:..10",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.score <= $1 )",
				"posts.score DESC LIMIT $2 OFFSET $3",
				[]interface{}{"10", 10, 0},
			},
		},
		{
			name: "rangable field less than",
			search: models.Search{
				Text:     "score:1..",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.score >= $1 )",
				"posts.score DESC LIMIT $2 OFFSET $3",
				[]interface{}{"1", 10, 0},
			},
		},
		{
			name: "parseable field",
			search: models.Search{
				Text:     "createdAt:2020-01-01..2020-01-02",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.created_at BETWEEN $1 AND $2 )",
				"posts.score DESC LIMIT $3 OFFSET $4",
				[]interface{}{"2020-01-01", "2020-01-02", 10, 0},
			},
		},
		{
			name: "sorting",
			search: models.Search{
				Text:     "tag:cat sort-asc:score",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.tags @> $1 )",
				"posts.score ASC LIMIT $2 OFFSET $3",
				[]interface{}{"cat", 10, 0},
			},
		},
		{
			name: "multiple sorting",
			search: models.Search{
				Text:     "tag:cat sort-asc:score sort-desc:fav",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.tags @> $1 )",
				"posts.score ASC, posts.fav_count DESC LIMIT $2 OFFSET $3",
				[]interface{}{"cat", 10, 0},
			},
		},
		{
			name: "sorting with not allowed order",
			search: models.Search{
				Text:     "tag:cat sort-asc:fav",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"( posts.tags @> $1 )",
				"posts.fav_count DESC LIMIT $2 OFFSET $3",
				[]interface{}{"cat", 10, 0},
			},
		},
		{
			name: "sorting with not existing field",
			search: models.Search{
				Text:     "sort-asc:non-existing",
				Page:     1,
				PageSize: 10,
			},
			want: []interface{}{
				"1 = 1",
				"posts.score DESC LIMIT $1 OFFSET $2",
				[]interface{}{10, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := parser.ParseSearch(tt.search)

			assert.NoError(t, err)
			assert.Equal(t, tt.want[0], r.WhereQuery)
			assert.Equal(t, tt.want[1], r.SortQuery)
			assert.Equal(t, tt.want[2], append(r.WhereArgs, r.PaginationArgs...))
		})
	}
}
