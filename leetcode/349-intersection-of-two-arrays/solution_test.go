package intersectionoftwoarrays

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{1}, []int{2}, []int{}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{1, 1, 1, 1}, []int{1, 1}, []int{1}},
		{[]int{5, 5, 5}, []int{5}, []int{5}},
		{[]int{1, 2, 3}, []int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{2, 4}, []int{2, 4}},
		{[]int{0, 1, 2}, []int{2, 0}, []int{0, 2}},
		{[]int{1000}, []int{1000}, []int{1000}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1, 0}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		got := intersection(slices.Clone(tt.nums1), slices.Clone(tt.nums2))
		want := slices.Clone(tt.want)
		slices.Sort(got)
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Errorf("intersection(%v, %v) = %v, want %v (any order)", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}
