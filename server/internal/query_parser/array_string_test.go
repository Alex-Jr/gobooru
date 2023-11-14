package query_parser_test

import (
	"fmt"
	"gobooru/internal/query_parser"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestArrayStringParserFn(t *testing.T) {
	testCases := []struct {
		input         interface{}
		expectedValue interface{}
	}{
		{true, &pq.BoolArray{true}},
		{int32(1), &pq.Int32Array{1}},
		{int64(1), &pq.Int64Array{1}},
		{"a", &pq.StringArray{"a"}},
		{float32(1.1), &pq.Float32Array{1.1}},
		{float64(1.1), &pq.Float64Array{1.1}},
		{[]interface{}{"a", "b", "c"}, pq.GenericArray{
			A: []interface{}{"a", "b", "c"},
		}},
		{[]interface{}{1, 2, 3}, pq.GenericArray{
			A: []interface{}{1, 2, 3},
		}},
	}

	for _, tc := range testCases {
		output := query_parser.ArrayParserFn(tc.input)
		if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", tc.expectedValue) {
			assert.Equal(t, tc.expectedValue, output)
		}
	}
}
