package pascalstriangle

import (
	"slices"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		num  int
		want [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{3, [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{4, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{6, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}}},
	}

	for _, tt := range tests {
		got := generate(tt.num)
		if !slices.EqualFunc(got, tt.want, slices.Equal) {
			t.Errorf("generate(%d) = %v, want %v", tt.num, got, tt.want)
		}
	}
}
