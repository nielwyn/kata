package findallduplicatesinanarray

import (
	"slices"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example 1", []int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{"example 2", []int{1, 1, 2}, []int{1}},
		{"example 3", []int{1}, []int{}},
		{"no duplicates", []int{1, 2, 3, 4}, []int{}},
		{"no duplicates, shuffled", []int{3, 1, 4, 2}, []int{}},
		{"every value duplicated", []int{2, 2, 1, 1, 3, 3}, []int{1, 2, 3}},
		{"duplicates in descending order", []int{3, 3, 1, 1}, []int{1, 3}},
		{"duplicate pair adjacent", []int{1, 2, 2, 4}, []int{2}},
		{"duplicate split to both ends", []int{4, 1, 2, 4}, []int{4}},
		{"lowest value duplicated", []int{5, 1, 3, 1, 4}, []int{1}},
		{"highest value duplicated", []int{5, 1, 3, 5, 4}, []int{5}},
		{"two elements", []int{2, 2}, []int{2}},
	}

	for _, tt := range tests {
		in := slices.Clone(tt.nums)
		got := findDuplicates(tt.nums)
		slices.Sort(got)
		if !slices.Equal(got, tt.want) {
			t.Errorf("%s: findDuplicates(%v) = %v, want %v", tt.name, in, got, tt.want)
		}
	}
}
