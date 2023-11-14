package query_parser

import (
	"fmt"
	"gobooru/internal/models"
	"regexp"
	"strings"
)

type ParserFn func(interface{}) interface{}

type ParserConfig struct {
	WhereField        map[string]WhereField
	SortField         map[string]SortField
	DefaultWhereField string
	DefaultSortField  string
}

type WhereField struct {
	Operator string
	DBName   string
	Rangable bool
	ParserFn ParserFn
}

type SortField struct {
	Name         string
	DBName       string
	AllowASC     bool
	AllowDESC    bool
	DefaultOrder string
}

type parser struct {
	WhereField        map[string]WhereField
	SortField         map[string]SortField
	DefaultWhereField string
	DefaultSortField  string
}

type Parser interface {
	ParseSearch(search models.Search) (parserResponse, error)
}

func NewParser(config ParserConfig) Parser {
	return &parser{
		WhereField:        config.WhereField,
		SortField:         config.SortField,
		DefaultWhereField: config.DefaultWhereField,
		DefaultSortField:  config.DefaultSortField,
	}
}

var spaceRegex = regexp.MustCompile(`\s+`)
var colonRegex = regexp.MustCompile(`:`)
var commaRegex = regexp.MustCompile(`,`)
var twoDots = regexp.MustCompile(`\.\.`)

// var dashRegex = regexp.MustCompile(`-`)

type parserResponse struct {
	WhereQuery     string
	SortQuery      string
	WhereArgs      []interface{}
	PaginationArgs []interface{}
}

// parses a string like 'tag:a,b' into ['tag:a', 'tag:b']
func commaSplit(s string) []string {
	k := colonRegex.Split(s, -1)

	// no prefix
	if len(k) == 1 {
		return commaRegex.Split(s, -1)
	}

	prefix := k[0]
	parts := strings.Split(s[len(prefix)+1:], ",")
	result := make([]string, len(parts))

	for i, part := range parts {
		result[i] = fmt.Sprintf("%s:%s", prefix, part)
	}

	return result
}

func (p parser) ParseSearch(search models.Search) (parserResponse, error) {
	paginationArgs := []interface{}{
		search.PageSize,
		(search.Page - 1) * search.PageSize,
	}

	whereQuery := make([]string, 0)
	sortQuery := make([]string, 0)

	whereArgs := make([]interface{}, 0)

	valueIndex := 0

	// split words by spaces
	// a,b c:1 => ['a,b' 'c:1']
	for _, word := range spaceRegex.Split(search.Text, -1) {

		if len(word) == 0 {
			continue
		}

		where := make([]string, 0)
		sort := make([]string, 0)

		// split term by ,
		// a,b => ['a' 'b']
		// c:1 => ['c:1']
		// c:1,2 => ['c:1' '2']
		for _, condition := range commaSplit(word) {
			if len(condition) == 0 {
				continue
			}

			// split condition by :
			// ['a'] => ['a']
			// ['c:1'] => ['c', 1]
			kv := colonRegex.Split(condition, -1)
			key := ""
			value := ""

			if len(kv) > 2 || len(kv) == 0 {
				continue
			} else if len(kv) == 2 {
				key = kv[0]
				value = kv[1]
			} else {
				key = p.DefaultWhereField
				value = kv[0]
			}

			if strings.HasPrefix(key, "sort") {
				field, exist := p.SortField[value]

				if !exist {
					continue
				}

				var order string

				order = field.DefaultOrder
				if key == "sort-asc" && field.AllowASC {
					order = "ASC"
				} else if key == "sort-desc" && field.AllowDESC {
					order = "DESC"
				}

				sort = append(sort, fmt.Sprintf("%s %s", field.DBName, order))
				continue
			}

			field, exist := p.WhereField[key]
			if !exist {
				continue
			}

			if !field.Rangable {
				operator := field.Operator
				negation := ""

				if strings.HasPrefix(value, "-") {
					negation = "NOT "
					value = value[1:]
				}

				valueIndex += 1
				where = append(where, fmt.Sprintf("%s%s %s $%d", negation, field.DBName, operator, valueIndex))
				if field.ParserFn == nil {
					whereArgs = append(whereArgs, value)
				} else {
					whereArgs = append(whereArgs, field.ParserFn(value))
				}

				continue
			}

			rangeQuery := twoDots.Split(value, -1)
			rangeQueryLen := len(rangeQuery)

			if rangeQueryLen == 2 {
				valueIndex += 1
				if rangeQuery[0] == "" {
					whereArgs = append(whereArgs, rangeQuery[1])
					where = append(where, fmt.Sprintf("%s <= $%d", field.DBName, valueIndex))
				} else if rangeQuery[1] == "" {
					whereArgs = append(whereArgs, rangeQuery[0])
					where = append(where, fmt.Sprintf("%s >= $%d", field.DBName, valueIndex))
				} else {
					where = append(where, fmt.Sprintf("%s BETWEEN $%d AND $%d", field.DBName, valueIndex, valueIndex+1))
					valueIndex += 1

					whereArgs = append(whereArgs, rangeQuery[0])
					whereArgs = append(whereArgs, rangeQuery[1])
				}
			}
		}

		if len(where) >= 1 {
			whereQuery = append(whereQuery, fmt.Sprintf("( %s )", strings.Join(where, " OR ")))
		} else {
			sortQuery = append(sortQuery, sort...)
		}
	}

	// put default value here
	if len(sortQuery) == 0 {
		fieldConfig := p.SortField[p.DefaultSortField]

		sortQuery = append(sortQuery, fmt.Sprintf("%s %s", fieldConfig.DBName, fieldConfig.DefaultOrder))
	}

	if len(whereQuery) == 0 {
		whereQuery = append(whereQuery, "1 = 1")
	}

	return parserResponse{
		WhereQuery:     strings.Join(whereQuery, " AND "),
		SortQuery:      strings.Join(sortQuery, ", ") + fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(whereArgs)+1, len(whereArgs)+2),
		WhereArgs:      whereArgs,
		PaginationArgs: paginationArgs,
	}, nil
}
