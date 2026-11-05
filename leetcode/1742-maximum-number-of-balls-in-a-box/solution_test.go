package maximumnumberofballsinabox

import "testing"

func TestCountBalls(t *testing.T) {
	tests := []struct {
		lowLimit  int
		highLimit int
		want      int
	}{
		{1, 10, 2},
		{5, 15, 2},
		{19, 28, 2},
		{1, 1, 1},
		{100000, 100000, 1},
		{1, 2, 1},
		{1, 9, 1},
		{9, 10, 1},
		{91, 100, 1},
		{1, 20, 3},
		{1, 99, 10},
		{1, 100000, 6000},
		{99999, 100000, 1},
		{50000, 100000, 3256},
	}

	for _, tt := range tests {
		got := countBalls(tt.lowLimit, tt.highLimit)
		if got != tt.want {
			t.Errorf("countBalls(%d, %d) = %d, want %d", tt.lowLimit, tt.highLimit, got, tt.want)
		}
	}
}
