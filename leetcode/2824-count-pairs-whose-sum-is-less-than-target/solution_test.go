package countpairswhosesumislessthantarget

import "testing"

func TestCountPairs(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 1, 2, 3, 1}, 2, 3},
		{[]int{-6, 2, 5, -2, -7, -1, 3}, -2, 10},
		{[]int{5}, 10, 0},
		{[]int{1, 2}, 4, 1},
		{[]int{1, 2}, 3, 0},
		{[]int{1, 1, 1}, 3, 3},
		{[]int{10, 20, 30}, 5, 0},
		{[]int{-50, -50}, -99, 1},
		{[]int{-50, -50}, -100, 0},
	}

	for _, tt := range tests {
		got := countPairs(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("countPairs(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
