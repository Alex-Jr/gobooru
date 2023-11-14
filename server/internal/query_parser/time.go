package query_parser

import "time"

func TimeParserFn(i interface{}) interface{} {
	parsed, err := time.Parse(time.RFC3339Nano, i.(string))

	if err != nil {
		return time.Time{}
	}

	return parsed
}
