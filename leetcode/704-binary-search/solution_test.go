package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{5}, -1, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{[]int{-1, 0, 3, 5, 9, 12}, -2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 13, -1},
	}

	for _, tt := range tests {
		got := search(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
