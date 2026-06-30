package findtheduplicatenumber

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
		{[]int{1, 1}, 1},
		{[]int{2, 1, 2}, 2},
		{[]int{1, 2, 3, 4, 5, 5}, 5},
		{[]int{1, 2, 3, 4, 4, 5}, 4},
	}

	for _, tt := range tests {
		got := findDuplicate(tt.nums)
		if got != tt.expect {
			t.Errorf("findDuplicate(%v) = %v, want %v", tt.nums, got, tt.expect)
		}
	}
}
