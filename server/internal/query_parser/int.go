package query_parser

import "strconv"

func IntParserFn(value interface{}) interface{} {
	s, err := strconv.Atoi(value.(string))

	if err != nil {
		return 0
	}

	return s
}
