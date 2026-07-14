package longestsubsequencewithlimitedsum

import (
	"slices"
	"testing"
)

func TestAnswerQueries(t *testing.T) {
	tests := []struct {
		nums    []int
		queries []int
		want    []int
	}{
		{[]int{4, 5, 2, 1}, []int{3, 10, 21}, []int{2, 3, 4}},
		{[]int{2, 3, 4, 5}, []int{1}, []int{0}},
		{[]int{5}, []int{5}, []int{1}},
		{[]int{5}, []int{4}, []int{0}},
		{[]int{1, 1, 1, 1}, []int{2, 4}, []int{2, 4}},
		{[]int{4, 5, 2, 1}, []int{12}, []int{4}},
		{[]int{3, 1, 2}, []int{7, 2, 6, 4}, []int{3, 1, 3, 2}},
		{[]int{1000000}, []int{1000000}, []int{1}},
	}

	for _, tt := range tests {
		got := answerQueries(tt.nums, tt.queries)
		if !slices.Equal(got, tt.want) {
			t.Errorf("answerQueries(%v, %v) = %v, want %v", tt.nums, tt.queries, got, tt.want)
		}
	}
}
