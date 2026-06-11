package degreeofarray

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
		{[]int{1}, 1},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 2, 3}, 1},
		{[]int{1, 2, 2, 1}, 2},
	}

	for _, tt := range tests {
		got := findShortestSubArray(tt.nums)
		if got != tt.want {
			t.Errorf("findShortestSubArray(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
