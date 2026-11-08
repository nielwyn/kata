package findallduplicatesinanarray

import (
	"slices"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 1, 2}, []int{1}},
		{[]int{1}, []int{}},
		{[]int{1, 2, 3, 4}, []int{}},
		{[]int{3, 1, 4, 2}, []int{}},
		{[]int{2, 2, 1, 1, 3, 3}, []int{1, 2, 3}},
		{[]int{3, 3, 1, 1}, []int{1, 3}},
		{[]int{1, 2, 2, 4}, []int{2}},
		{[]int{4, 1, 2, 4}, []int{4}},
		{[]int{5, 1, 3, 1, 4}, []int{1}},
		{[]int{5, 1, 3, 5, 4}, []int{5}},
		{[]int{2, 2}, []int{2}},
	}

	for _, tt := range tests {
		in := slices.Clone(tt.nums)
		got := findDuplicates(tt.nums)
		slices.Sort(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("findDuplicates(%v) = %v, want %v", in, got, tt.want)
		}
	}
}
