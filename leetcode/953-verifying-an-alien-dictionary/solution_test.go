package verifyinganaliendictionary

import "testing"

const (
	normal   = "abcdefghijklmnopqrstuvwxyz"
	reversed = "zyxwvutsrqponmlkjihgfedcba"
)

func TestIsAlienSorted(t *testing.T) {
	tests := []struct {
		words []string
		order string
		want  bool
	}{
		{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
		{[]string{"apple", "app"}, normal, false},
		{[]string{"app", "apple"}, normal, true},
		{[]string{"leetcode"}, normal, true},
		{[]string{"a"}, reversed, true},
		{[]string{"abc", "abc"}, normal, true},
		{[]string{"zz", "zz"}, reversed, true},
		{[]string{"a", "b", "c", "d"}, normal, true},
		{[]string{"a", "b", "c", "d"}, reversed, false},
		{[]string{"zyx", "abc"}, reversed, true},
		{[]string{"abc", "zyx"}, reversed, false},
		{[]string{"aa", "ab", "aa"}, normal, false},
		{[]string{"a", "b", "c", "b"}, normal, false},
		{[]string{"abcde", "abcdf"}, normal, true},
		{[]string{"abcdf", "abcde"}, normal, false},
		{[]string{"az", "ba"}, normal, true},
		{[]string{"ba", "az"}, normal, false},
		{[]string{"zebra", "apple"}, reversed, true},
		{[]string{"a", "aa", "aaa"}, normal, true},
		{[]string{"aaa", "aa", "a"}, normal, false},
		{[]string{"ab", "ab", "ac"}, normal, true},
		{[]string{"aaaaaaaaaaaaaaaaaaab", "aaaaaaaaaaaaaaaaaaac"}, normal, true},
		{[]string{"aaaaaaaaaaaaaaaaaaac", "aaaaaaaaaaaaaaaaaaab"}, normal, false},
		{[]string{"hlabc", "hlabd"}, "hlabcdefgijkmnopqrstuvwxyz", true},
		{[]string{"abcdefghijklm", "nopqrstuvwxyz"}, normal, true},
	}

	for _, tt := range tests {
		got := isAlienSorted(tt.words, tt.order)
		if got != tt.want {
			t.Errorf("isAlienSorted(%q, %q) = %t, want %t", tt.words, tt.order, got, tt.want)
		}
	}
}
