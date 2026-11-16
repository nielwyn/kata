package validanagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "a", true},
		{"a", "b", false},
		{"ab", "a", false},
		{"a", "ab", false},
		{"aacc", "ccac", false},
		{"listen", "silent", true},
		{"aabbcc", "ccbbaa", true},
		{"abcd", "abcd", true},
		{"aaz", "zza", false},
		{"xxyyz", "zyxyx", true},
	}

	for _, tt := range tests {
		got := isAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}
