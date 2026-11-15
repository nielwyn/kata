package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 1},
		{[]int{2, 1, 1}, 1},
		{[]int{1, 2, 1}, 1},
		{[]int{6, 5, 5}, 5},
		{[]int{8, 8, 7, 7, 7}, 7},
		{[]int{3, 3, 4, 2, 4, 4, 2, 4, 4}, 4},
		{[]int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5}, 5},
		{[]int{5, 5, 5, 5, 5, 5, 1, 2, 3, 4}, 5},
		{[]int{-1, -1, -1, 2, 2}, -1},
		{[]int{0, 0, 0, 1, 2}, 0},
		{[]int{1000000000, -1000000000, 1000000000}, 1000000000},
		{[]int{-1000000000, -1000000000, 1000000000, 1000000000, -1000000000}, -1000000000},
	}

	for _, tt := range tests {
		got := majorityElement(tt.nums)
		if got != tt.want {
			t.Errorf("majorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
