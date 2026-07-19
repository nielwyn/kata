package checkifthesentenceispangram

import "testing"

func TestCheckIfPangram(t *testing.T) {
	tests := []struct {
		sentence string
		want     bool
	}{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"abcdefghijklmnopqrstuvwxy", false},
		{"a", false},
		{"aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz", true},
		{"thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog", true},
	}

	for _, tt := range tests {
		got := checkIfPangram(tt.sentence)
		if got != tt.want {
			t.Errorf("checkIfPangram(%q) = %v, want %v", tt.sentence, got, tt.want)
		}
	}
}
