package query_parser

import "fmt"

func StartsWithParserFn(value interface{}) interface{} {
	return fmt.Sprintf("%s%%", value.(string))
}
