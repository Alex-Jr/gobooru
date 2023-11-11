package query_parser

import (
	"errors"
	"fmt"
	"gobooru/internal/models"
	"regexp"
	"strconv"
	"strings"
)

type ParserFn func(interface{}) interface{}

type ParserConfig struct {
	WhereField   map[string]WhereField
	SortField    map[string]SortField
	DefaultWhere string
	DefaultSort  string
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
	WhereField   map[string]WhereField
	SortField    map[string]SortField
	Default      string
	DefaultSort  string
	DefaultQuery string
}

type Parser interface {
	Parse(search models.Search) (string, string, []interface{}, []interface{}, error)
}

func NewParser(config ParserConfig) Parser {
	return &parser{
		WhereField:  config.WhereField,
		SortField:   config.SortField,
		Default:     config.DefaultWhere,
		DefaultSort: config.DefaultSort,
	}
}

var spaceRegex = regexp.MustCompile(`\s+`)
var colonRegex = regexp.MustCompile(`:`)
var commaRegex = regexp.MustCompile(`,`)
var twoDots = regexp.MustCompile(`\.\.`)

// var dashRegex = regexp.MustCompile(`-`)

func (p parser) Parse(search models.Search) (string, string, []interface{}, []interface{}, error) {
	metaArgs := []interface{}{
		search.PageSize,
		(search.Page - 1) * search.PageSize,
	}

	whereQuery := " 1 = 1 "
	sortQuery := " "

	whereArgs := make([]interface{}, 0)

	valueIndex := 0

	// split terms by spaces
	// a,b c:1 => ['a,b' 'c:1']
	for _, term := range spaceRegex.Split(search.Text, -1) {
		// split term by ,
		// a,b => ['a' 'b']
		// c:1 => ['c:1']
		conditions := commaRegex.Split(term, -1)

		where := make([]string, 0)
		sort := ""
		for _, condition := range conditions {
			// split condition by :
			// ['a'] => ['a']
			// ['c:1'] => ['c', 1]
			keyValue := colonRegex.Split(condition, -1)

			if len(keyValue) > 2 {
				return "", "", []interface{}{}, []interface{}{}, errors.New(": inside value")
			} else if len(keyValue) == 1 {
				// anonymous will be transformed to a named condition using parser default
				keyValue = append([]string{p.Default}, keyValue...)
			}

			key := keyValue[0]
			value := keyValue[1]

			if strings.HasPrefix(key, "sort") {
				fieldConfig, exist := p.SortField[value]

				if !exist {
					return "", "", []interface{}{}, []interface{}{}, errors.New("sort by unknown field")
				}

				var order string

				order = fieldConfig.DefaultOrder
				if key == "sort-asc" && fieldConfig.AllowASC {
					order = "ASC"
				} else if key == "sort-asc" && fieldConfig.AllowDESC {
					order = "DESC"
				}

				sort = fmt.Sprintf("%s %s", fieldConfig.DBName, order)
				continue
			}

			fieldConfig, exist := p.WhereField[key]
			if !exist {
				return "", "", []interface{}{}, []interface{}{}, errors.New("search by unknown field")
			}

			if !fieldConfig.Rangable {
				operator := fieldConfig.Operator
				negation := ""

				if strings.HasPrefix(value, "-") {
					negation = "NOT"
					value = value[1:]
				}

				valueIndex += 1
				where = append(where, fmt.Sprintf("%s %s %s $%d", negation, fieldConfig.DBName, operator, valueIndex))
				if fieldConfig.ParserFn == nil {
					whereArgs = append(whereArgs, value)
				} else {
					whereArgs = append(whereArgs, fieldConfig.ParserFn(value))
				}

				continue
			}

			rangeQuery := twoDots.Split(value, -1)
			rangeQueryLen := len(rangeQuery)

			if rangeQueryLen == 2 {
				valueIndex += 1
				if rangeQuery[0] == "" {
					whereArgs = append(whereArgs, rangeQuery[1])
					where = append(where, fmt.Sprintf("%s <= $%d", fieldConfig.DBName, valueIndex))
				} else if rangeQuery[1] == "" {
					whereArgs = append(whereArgs, rangeQuery[0])
					where = append(where, fmt.Sprintf("%s >= $%d", fieldConfig.DBName, valueIndex))
				} else {
					where = append(where, fmt.Sprintf("%s BETWEEN $%d AND $%d", fieldConfig.DBName, valueIndex, valueIndex+1))
					valueIndex += 1

					whereArgs = append(whereArgs, rangeQuery[0])
					whereArgs = append(whereArgs, rangeQuery[1])
				}
			}
		}

		if len(where) > 1 {
			whereQuery += fmt.Sprintf(" AND ( %s )", strings.Join(where, " OR "))
		} else if len(where) == 1 {
			whereQuery += " AND " + strings.Join(where, "")
		} else {
			sortQuery += sort + ", "
		}
	}

	// put default value here
	if len(sortQuery) == 1 {
		fieldConfig := p.SortField[p.DefaultSort]

		sortQuery = fmt.Sprintf(" %s %s, ", fieldConfig.DBName, fieldConfig.DefaultOrder)
	}

	if len(whereQuery) == 1 {
		whereQuery = p.DefaultQuery
	}

	sortQuery = sortQuery[:len(sortQuery)-2]

	sortQuery += `
		LIMIT $` + strconv.Itoa(len(whereArgs)+1) + `
		OFFSET $` + strconv.Itoa(len(whereArgs)+2)

	return whereQuery, sortQuery, whereArgs, metaArgs, nil
}
