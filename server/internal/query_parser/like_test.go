package query_parser_test

import (
	"gobooru/internal/query_parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikeParserFn(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected interface{}
	}{
		{
			input:    "cat",
			expected: "%cat%",
		},
		{
			input:    "cat%",
			expected: "%cat%%",
		},
		{
			input:    "%cat",
			expected: "%%cat%",
		},
		{
			input:    "%cat%",
			expected: "%%cat%%",
		},
		{
			input:    "cat%cat",
			expected: "%cat%cat%",
		},
		{
			input:    "123",
			expected: "%123%",
		},
	}

	for _, tc := range testCases {
		output := query_parser.LikeParserFn(tc.input)
		assert.Equal(t, tc.expected, output)
	}
}
