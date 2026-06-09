package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interview", "interact", "intercom"}, "inter"},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"abc", "abc"}, "abc"},
		{[]string{""}, ""},
	}

	for _, tt := range tests {
		got := longestCommonPrefix(tt.strs)
		if got != tt.want {
			t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
		}
	}
}
