package missingnumber

import (
	"slices"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
		{[]int{1}, 0},
		{[]int{1, 0, 3}, 2},
		{[]int{0, 1, 2}, 3},
		{[]int{1, 2, 3}, 0},
		{[]int{2, 0}, 1},
		{[]int{4, 1, 3, 0, 5, 2}, 6},
		{[]int{6, 5, 4, 3, 2, 1}, 0},
	}

	for _, tt := range tests {
		in := slices.Clone(tt.nums)
		got := missingNumber(tt.nums)
		if got != tt.want {
			t.Errorf("missingNumber(%v) = %d, want %d", in, got, tt.want)
		}
	}
}
