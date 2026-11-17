package majorityelement2

import (
	"slices"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{3, 2, 3}, []int{3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3, 4}, []int{}},
		{[]int{2, 2}, []int{2}},
		{[]int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
		{[]int{3, 3, 4}, []int{3}},
		{[]int{4, 3, 3}, []int{3}},
		{[]int{3, 4, 3}, []int{3}},
		{[]int{1, 1, 2, 2, 3, 3}, []int{}},
		{[]int{1, 1, 1, 2, 3, 4, 5, 6}, []int{1}},
		{[]int{0, 0, 0}, []int{0}},
		{[]int{-1, -1, -1, 2, 2}, []int{-1, 2}},
		{[]int{-1, 2, 3}, []int{}},
		{[]int{1000000000, 1000000000, -1000000000, -1000000000, 7}, []int{-1000000000, 1000000000}},
		{[]int{5, 5, 5, 5, 1, 2, 3, 4, 6, 7}, []int{5}},
		{[]int{1, 2, 3, 4, 5, 5, 5, 5, 6, 7}, []int{5}},
		{[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, []int{4}},
		{[]int{7, 7, 7, 7, 7, 7, 7}, []int{7}},
		{[]int{1, 2, 3, 1, 2, 3, 1, 2, 3}, []int{}},
		{[]int{8, 8, 8, 8, 9, 9, 9, 9, 1}, []int{8, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}},
	}

	for _, tt := range tests {
		in := slices.Clone(tt.nums)
		got := majorityElement(tt.nums)
		want := slices.Clone(tt.want)
		slices.Sort(got)
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Errorf("majorityElement(%v) = %v, want %v (any order)", in, got, tt.want)
		}
	}
}
