package findsmallestlettergreaterthantarget

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'}, 'e', 'n'},
		{[]byte{'a', 'b'}, 'z', 'a'},
	}

	for _, tt := range tests {
		got := nextGreatestLetter(tt.letters, tt.target)
		if got != tt.want {
			t.Errorf("nextGreatestLetter(%s, %q) = %q, want %q", tt.letters, tt.target, got, tt.want)
		}
	}
}
