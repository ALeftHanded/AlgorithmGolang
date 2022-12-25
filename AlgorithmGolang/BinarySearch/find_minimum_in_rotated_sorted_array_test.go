package BinarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "normal test even elements",
			args: args{
				nums: []int{4, 5, 6, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "normal test odd elements",
			args: args{
				nums: []int{4, 5, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "test single element",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "test ascending array",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 1,
		},
		{
			// ! 需要cover这个用例
			name: "test two elements and descending array",
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FindMinimumInRotatedSortedArray(tt.args.nums), "FindMinimumInRotatedSortedArray(%v)", tt.args.nums)
		})
	}
}
