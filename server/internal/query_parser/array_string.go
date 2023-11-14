package query_parser

import (
	"github.com/lib/pq"
)

func ArrayParserFn(i interface{}) interface{} {
	switch v := i.(type) {
	case bool:
		return pq.Array([]bool{v})
	case int32:
		return pq.Array([]int32{v})
	case int64:
		return pq.Array([]int64{v})
	case float32:
		return pq.Array([]float32{v})
	case float64:
		return pq.Array([]float64{v})
	case string:
		return pq.Array([]string{v})
	default:
		return pq.Array(v)
	}
}
