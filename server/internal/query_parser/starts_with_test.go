package query_parser_test

import (
	"gobooru/internal/query_parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsWithParserFn(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	testCases := []testCase{{
		input:    "test",
		expected: "test%",
	}, {
		input:    "test%",
		expected: "test%%",
	}, {
		input:    "test%%",
		expected: "test%%%",
	}, {
		input:    "",
		expected: "%",
	}}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := query_parser.StartsWithParserFn(tc.input)

			assert.Equal(t, tc.expected, result)
		})
	}
}
