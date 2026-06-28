package firstbadversion

import "testing"

var firstBad int

func init() {
	isBadVersion = func(version int) bool {
		return version >= firstBad
	}
}

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		n      int
		bad    int
		expect int
	}{
		{5, 4, 4},
		{1, 1, 1},
		{10, 1, 1},
		{10, 10, 10},
		{10, 5, 5},
		{100, 50, 50},
	}

	for _, tt := range tests {
		firstBad = tt.bad
		got := firstBadVersion(tt.n)
		if got != tt.expect {
			t.Errorf("firstBadVersion(%v) with bad=%v = %v, want %v", tt.n, tt.bad, got, tt.expect)
		}
	}
}
