package findthedifference

import (
	"testing"
)

func TestFindTheDifference(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want byte
	}{
		{"abcd", "abcde", 'e'},
		{"", "y", 'y'},
		{"a", "aa", 'a'},
		{"ae", "aea", 'a'},
	}

	for _, tt := range tests {
		got := findTheDifference(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("findTheDifference(%q, %q) = %c, want %c", tt.s, tt.t, got, tt.want)
		}
	}
}
