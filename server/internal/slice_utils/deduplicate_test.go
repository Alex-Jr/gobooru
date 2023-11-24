package slice_utils_test

import (
	"gobooru/internal/slice_utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeduplicateInt(t *testing.T) {
	type args struct {
		a []int
	}

	type want struct {
		result []int
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty",
			args: args{
				a: []int{},
			},
			want: want{
				result: []int{},
			},
		},
		{
			name: "no deduplication",
			args: args{
				a: []int{1, 2, 3},
			},
			want: want{
				result: []int{1, 2, 3},
			},
		},
		{
			name: "deduplication",
			args: args{
				a: []int{1, 2, 3, 1, 2, 3},
			},
			want: want{
				result: []int{1, 2, 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slice_utils.Deduplicate(tc.args.a)

			require.Equal(t, len(tc.want.result), len(result))

			seen := make(map[int]bool, len(tc.want.result))

			for _, item := range result {
				seen[item] = true
			}

			for _, item := range tc.want.result {
				assert.True(t, seen[item])
			}
		})
	}
}

func TestDeduplicateString(t *testing.T) {
	type args struct {
		a []string
	}

	type want struct {
		result []string
	}

	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty",
			args: args{
				a: []string{},
			},
			want: want{
				result: []string{},
			},
		},
		{
			name: "no deduplication",
			args: args{
				a: []string{"a", "b", "c"},
			},
			want: want{
				result: []string{"a", "b", "c"},
			},
		},
		{
			name: "deduplication",
			args: args{
				a: []string{"a", "b", "c", "a", "b", "c"},
			},
			want: want{
				result: []string{"a", "b", "c"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slice_utils.Deduplicate(tc.args.a)

			seen := make(map[string]bool, len(tc.want.result))

			require.Equal(t, len(tc.want.result), len(result))

			for _, item := range result {
				seen[item] = true
			}

			for _, item := range tc.want.result {
				assert.True(t, seen[item])
			}
		})
	}
}
