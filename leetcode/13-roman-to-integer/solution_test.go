package romantointeger

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"IV", 4},
		{"IX", 9},
		{"XL", 40},
		{"XC", 90},
		{"CD", 400},
		{"CM", 900},
		{"M", 1000},
	}

	for _, tt := range tests {
		got := romanToInt(tt.s)
		if got != tt.want {
			t.Errorf("romanToInt(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}
