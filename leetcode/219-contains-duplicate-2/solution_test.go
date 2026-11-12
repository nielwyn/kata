package containsduplicate2

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1}, 0, false},
		{[]int{1}, 1, false},
		{[]int{1, 1}, 0, false},
		{[]int{1, 1}, 1, true},
		{[]int{1, 2, 1}, 1, false},
		{[]int{1, 2, 1}, 2, true},
		{[]int{1, 2, 3, 4, 5}, 100000, false},
		{[]int{1, 2, 3, 4, 1}, 100000, true},
		{[]int{0, 0}, 1, true},
		{[]int{-1000000000, 1000000000, -1000000000}, 2, true},
		{[]int{-1000000000, 1000000000, -1000000000}, 1, false},
		{[]int{1, 2, 2, 3, 1}, 1, true},
		{[]int{99, 99}, 2, true},
		{[]int{4, 1, 2, 3, 1, 5}, 3, true},
		{[]int{4, 1, 2, 3, 1, 5}, 2, false},
	}

	for _, tt := range tests {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("containsNearbyDuplicate(%v, %d) = %t, want %t", tt.nums, tt.k, got, tt.want)
		}
	}
}
