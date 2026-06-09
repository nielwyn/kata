package palindromenumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{121, true},
		{1221, true},
		{0, true},
		{1, true},
		{11, true},
		{-121, false},
		{10, false},
		{123, false},
		{-1, false},
	}

	for _, tt := range tests {
		got := isPalindrome(tt.x)
		if got != tt.want {
			t.Errorf("isPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
		}
	}
}
