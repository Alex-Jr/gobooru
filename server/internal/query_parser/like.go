package query_parser

import "fmt"

func LikeParserFn(i interface{}) interface{} {
	return fmt.Sprintf("%%%s%%", i)
}
