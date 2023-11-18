package slice_utils_test

import (
	"gobooru/internal/slice_utils"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDifference(t *testing.T) {
	type args struct {
		a []string
		b []string
	}

	type want struct {
		expected []string
	}

	testsCase := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Difference of two empty slices",
			args: args{a: []string{}, b: []string{}},
			want: want{expected: []string{}},
		},
		{
			name: "Difference of two slices with common elements",
			args: args{a: []string{"apple", "banana", "cherry"}, b: []string{"banana", "cherry", "date"}},
			want: want{expected: []string{"apple"}},
		},
		{
			name: "Difference of two slices with no common elements",
			args: args{a: []string{"apple", "banana", "cherry"}, b: []string{"date", "eggplant", "fig"}},
			want: want{expected: []string{"apple", "banana", "cherry"}},
		},
	}

	for _, tc := range testsCase {
		t.Run(tc.name, func(t *testing.T) {
			result := slice_utils.Difference(tc.args.a, tc.args.b)

			assert.DeepEqual(t, result, tc.want.expected)
		})
	}
}
