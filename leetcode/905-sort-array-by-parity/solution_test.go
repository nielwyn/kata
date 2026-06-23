package sortarraybyparity

import (
	"slices"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := [][]int{
		{3, 1, 2, 4},
		{0},
		{1, 2, 4, 3},
		{2, 4, 6, 8},
		{1, 3, 5, 7},
		{0, 1},
	}

	for _, nums := range tests {
		got := sortArrayByParity(slices.Clone(nums))

		if !isPermutation(got, nums) {
			t.Errorf("sortArrayByParity(%v) = %v, not a permutation of input", nums, got)
			continue
		}
		if !isPartitionedByParity(got) {
			t.Errorf("sortArrayByParity(%v) = %v, evens are not all before odds", nums, got)
		}
	}
}

// isPartitionedByParity reports whether all even numbers appear before all odd numbers.
func isPartitionedByParity(nums []int) bool {
	seenOdd := false
	for _, n := range nums {
		if n%2 != 0 {
			seenOdd = true
		} else if seenOdd {
			return false
		}
	}
	return true
}

func isPermutation(a, b []int) bool {
	a, b = slices.Clone(a), slices.Clone(b)
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
