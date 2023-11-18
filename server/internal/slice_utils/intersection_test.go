package slice_utils

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{
			name:     "Intersection of two empty slices",
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
		{
			name:     "Intersection of two slices with common elements",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{4, 5, 6, 7, 8},
			expected: []int{4, 5},
		},
		{
			name:     "Intersection of two slices with no common elements",
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Intersection(tc.a, tc.b)

			assert.DeepEqual(t, result, tc.expected)
		})
	}
}
