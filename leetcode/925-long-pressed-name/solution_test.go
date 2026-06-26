package longpressedname

import "testing"

func TestIsLongPressedName(t *testing.T) {
	tests := []struct {
		name   string
		typed  string
		expect bool
	}{
		{"alex", "aaleex", true},
		{"alex", "alex", true},
		{"saeed", "ssaaeedd", true},
		{"a", "aaaa", true},
		{"ab", "aab", true},
		{"ab", "abb", true},
		{"aaa", "aaaa", true},

		{"saeed", "ssaaedd", false},
		{"alex", "alexd", false},
		{"alex", "ale", false},
		{"abc", "abcd", false},
		{"aa", "a", false},
		{"ab", "ba", false},
		{"a", "b", false},
		{"saeed", "saeedhhhh", false},
	}

	for _, tt := range tests {
		got := isLongPressedName(tt.name, tt.typed)
		if got != tt.expect {
			t.Errorf("isLongPressedName(%q, %q) = %v, want %v", tt.name, tt.typed, got, tt.expect)
		}
	}
}
