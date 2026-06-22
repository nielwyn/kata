package reversestring

import (
	"slices"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
		{[]byte("a"), []byte("a")},
		{[]byte("ab"), []byte("ba")},
	}

	for _, tt := range tests {
		s := slices.Clone(tt.s)
		reverseString(s)
		if !slices.Equal(s, tt.want) {
			t.Errorf("reverseString(%q) = %q, want %q", tt.s, s, tt.want)
		}
	}
}
