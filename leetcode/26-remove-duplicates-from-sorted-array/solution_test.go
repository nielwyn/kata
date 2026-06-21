package removeduplicatesfromsortedarray

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		wantK    int
		wantNums []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 1}, 1, []int{1}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		nums := slices.Clone(tt.nums)
		gotK := removeDuplicates(nums)
		if gotK != tt.wantK {
			t.Errorf("removeDuplicates(%v) = %d, want %d", tt.nums, gotK, tt.wantK)
			continue
		}
		if !slices.Equal(nums[:gotK], tt.wantNums) {
			t.Errorf("removeDuplicates(%v) nums[:%d] = %v, want %v", tt.nums, gotK, nums[:gotK], tt.wantNums)
		}
	}
}
