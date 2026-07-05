package findpeakelement

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 1},
		{1, 2, 1, 3, 5, 6, 4},
		{1},
		{1, 2},
		{2, 1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, 4, 3, 2, 1, 5, 6, 7},
	}

	for _, nums := range tests {
		got := findPeakElement(nums)
		if got < 0 || got >= len(nums) {
			t.Errorf("findPeakElement(%v) = %d, out of bounds", nums, got)
			continue
		}
		if !isPeak(nums, got) {
			t.Errorf("findPeakElement(%v) = %d, index %d is not a peak", nums, got, got)
		}
	}
}

// isPeak reports whether nums[idx] is strictly greater than its neighbors,
// treating out-of-bounds neighbors as -infinity.
func isPeak(nums []int, idx int) bool {
	if idx > 0 && nums[idx] <= nums[idx-1] {
		return false
	}
	if idx < len(nums)-1 && nums[idx] <= nums[idx+1] {
		return false
	}
	return true
}
