package arrangingcoins

import "testing"

func TestArrangeCoins(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 2},
		{8, 3},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{2147483647, 65535},
	}

	for _, tt := range tests {
		got := arrangeCoins(tt.n)
		if got != tt.want {
			t.Errorf("arrangeCoins(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
