package nrepeatedelementinsize2narray

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example 1", []int{1, 2, 3, 3}, 3},
		{"example 2", []int{2, 1, 2, 5, 3, 2}, 2},
		{"example 3", []int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
		{"adjacent at start", []int{4, 4, 1, 2}, 4},
		{"split to both ends", []int{7, 1, 2, 7}, 7},
		{"clustered at end", []int{1, 2, 3, 4, 1, 1}, 1},
		{"repeated value is zero", []int{0, 1, 2, 0}, 0},
		{"max element value", []int{10000, 1, 10000, 2}, 10000},
		{"alternating", []int{1, 2, 1, 3, 1, 4}, 1},
	}

	for _, tt := range tests {
		in := slices.Clone(tt.nums)
		got := repeatedNTimes(tt.nums)
		if got != tt.want {
			t.Errorf("%s: repeatedNTimes(%v) = %d, want %d", tt.name, in, got, tt.want)
		}
	}
}

// Upper bound of the constraints: n == 5000, so len(nums) == 10000.
func TestRepeatedNTimesMaxSize(t *testing.T) {
	const n = 5000
	const want = 0

	nums := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		nums = append(nums, want)
	}
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	r := rand.New(rand.NewPCG(1, 2))
	r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	if got := repeatedNTimes(nums); got != want {
		t.Errorf("repeatedNTimes(shuffled n=%d) = %d, want %d", n, got, want)
	}
}
