package assigncookies

import "testing"

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g      []int
		s      []int
		expect int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4}, []int{1, 2}, 2},
		{[]int{5, 6, 7}, []int{1, 2, 3}, 0},
		{[]int{3, 1, 2}, []int{2, 1, 3}, 3},
		{[]int{1}, []int{1}, 1},
		{[]int{2}, []int{1}, 0},
		{[]int{1, 2}, []int{2, 2}, 2},
		{[]int{1, 2, 3}, []int{}, 0},
		{[]int{1, 3}, []int{3, 2, 1}, 2},
	}

	for _, tt := range tests {
		got := findContentChildren(tt.g, tt.s)
		if got != tt.expect {
			t.Errorf("findContentChildren(%v, %v) = %v, want %v", tt.g, tt.s, got, tt.expect)
		}
	}
}
