package validperfectsquare

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		num    int
		expect bool
	}{
		{16, true},
		{14, false},
		{1, true},
		{4, true},
		{9, true},
		{25, true},
		{2, false},
		{3, false},
		{2147395600, true},  // 46340 * 46340, near max int32
		{2147483647, false}, // max int32, not a perfect square
	}

	for _, tt := range tests {
		got := isPerfectSquare(tt.num)
		if got != tt.expect {
			t.Errorf("isPerfectSquare(%v) = %v, want %v", tt.num, got, tt.expect)
		}
	}
}
