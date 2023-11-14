package query_parser_test

import (
	"gobooru/internal/query_parser"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeParserFn(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "Test case 1",
			input:    "2021-08-01T00:00:00Z",
			expected: time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Test case 2",
			input:    "random string",
			expected: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := query_parser.TimeParserFn(tt.input)
			assert.Equal(t, tt.expected, output, "Test case %s failed", tt.name)
		})
	}
}
