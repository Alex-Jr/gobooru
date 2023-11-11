package query_parser

import "time"

func TimeParserFn(i interface{}) interface{} {
	time, _ := time.Parse(time.RFC3339Nano, i.(string))
	return time
}
