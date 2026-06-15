package checkdistancesbetweensameetters

import "testing"

func TestCheckDistances(t *testing.T) {
	tests := []struct {
		s        string
		distance []int
		want     bool
	}{
		{"abaccb", d26(map[int]int{0: 1, 1: 3, 2: 0, 3: 5}), true},
		{"aa", d26(map[int]int{0: 0}), true},
		{"aa", d26(map[int]int{0: 1}), false},
		{"bccb", d26(map[int]int{1: 2, 2: 0}), true},
	}

	for _, tt := range tests {
		got := checkDistances(tt.s, tt.distance)
		if got != tt.want {
			t.Errorf("checkDistances(%q, %v) = %v, want %v", tt.s, tt.distance, got, tt.want)
		}
	}
}

func d26(overrides map[int]int) []int {
	d := make([]int, 26)
	for k, v := range overrides {
		d[k] = v
	}
	return d
}
