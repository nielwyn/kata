package longestharmonioussubsequence

import (
	"testing"
)

func TestFindLHS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{}, 0},
		{[]int{5}, 0},
		{[]int{1, 1, 1, 1}, 0},
		{[]int{1, 2}, 2},
		{[]int{-1, -2, -2, -1}, 4},
	}

	for _, tt := range tests {
		got := findLHS(tt.nums)
		if got != tt.want {
			t.Errorf("findLHS(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
