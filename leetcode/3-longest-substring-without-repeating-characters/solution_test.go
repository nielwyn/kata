package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"a", 1},
		{"au", 2},
		{"aab", 2},
		{"abba", 2},
		{"dvdf", 3},
		{"tmmzuxt", 5},
		{"ohvhjdml", 6},
		{"abcdefg", 7},
		{"a1!a1", 3},
	}

	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.s)
		if got != tt.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
