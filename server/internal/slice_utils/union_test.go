package slice_utils_test

import (
	"gobooru/internal/slice_utils"
	"testing"

	"gotest.tools/v3/assert"
)

func TestUnion(t *testing.T) {
	type args struct {
		a []int
		b []int
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
			name: "Union of two empty slices",
			args: args{
				a: []int{},
				b: []int{},
			},
			want: want{
				result: []int{},
			},
		},
		{
			name: "Union of two slices with common elements",
			args: args{
				a: []int{1, 2, 3},
				b: []int{3, 4, 5},
			},
			want: want{
				result: []int{1, 2, 3, 4, 5},
			},
		},
		{
			name: "Union of two slices with no common elements",
			args: args{
				a: []int{1, 2, 3},
				b: []int{4, 5, 6},
			},
			want: want{
				result: []int{1, 2, 3, 4, 5, 6},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slice_utils.Union(tc.args.a, tc.args.b)

			resultMap := make(map[int]struct{})

			for _, v := range result {
				resultMap[v] = struct{}{}
			}

			for _, v := range tc.want.result {
				_, ok := resultMap[v]
				assert.Assert(t, ok, "expected %v to be in result", v)
			}
		})
	}
}
