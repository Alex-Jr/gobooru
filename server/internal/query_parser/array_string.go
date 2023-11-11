package query_parser

import "fmt"

func ArrayStringParserFn(i interface{}) interface{} {
	return fmt.Sprintf("{%s}", i)
}
