package sqrtx

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
		{2, 1},
		{9, 3},
		{2147483647, 46340},
	}
	for _, tt := range tests {
		got := mySqrt(tt.x)
		if got != tt.want {
			t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.want)
		}
	}
}
