package vowelspellchecker

import (
	"slices"
	"testing"
)

func TestSpellchecker(t *testing.T) {
	tests := []struct {
		wordlist []string
		queries  []string
		want     []string
	}{
		{
			[]string{"KiTe", "kite", "hare", "Hare"},
			[]string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
		{[]string{"yellow"}, []string{"YellOw"}, []string{"yellow"}},
		{[]string{"yellow"}, []string{"yellow"}, []string{"yellow"}},
		{[]string{"Yellow"}, []string{"yellow"}, []string{"Yellow"}},
		{[]string{"YellOw"}, []string{"yollow", "yeellow", "yllw"}, []string{"YellOw", "", ""}},

		// exact match wins over an earlier case-insensitive match
		{[]string{"Kite", "kite"}, []string{"kite"}, []string{"kite"}},
		{[]string{"kite", "Kite"}, []string{"Kite"}, []string{"Kite"}},
		// first case-insensitive match wins, not the closest one
		{[]string{"KiTe", "kite", "KITE"}, []string{"kItE"}, []string{"KiTe"}},
		// capitalization beats a vowel match that appears earlier
		{[]string{"kate", "KiTe"}, []string{"kIte"}, []string{"KiTe"}},
		// first vowel match in wordlist order
		{[]string{"kate", "kite"}, []string{"kote", "kute"}, []string{"kate", "kate"}},
		{[]string{"kite", "kate"}, []string{"kote"}, []string{"kite"}},

		// vowel substitution never changes length or consonants
		{[]string{"hello"}, []string{"hallo", "hullo", "heallo", "hllo", "hellp"}, []string{"hello", "hello", "", "", ""}},
		{[]string{"YeLLow"}, []string{"yillow", "YILLOW"}, []string{"YeLLow", "YeLLow"}},

		// single character
		{[]string{"a"}, []string{"a", "A", "e", "E", "b"}, []string{"a", "a", "a", "a", ""}},
		{[]string{"B"}, []string{"b", "B", "a"}, []string{"B", "B", ""}},

		// all vowels / no vowels
		{[]string{"aeiou"}, []string{"uoiea", "AEIOU", "aeiob"}, []string{"aeiou", "aeiou", ""}},
		{[]string{"rhythm"}, []string{"rhythm", "RHYTHM", "rhythma"}, []string{"rhythm", "rhythm", ""}},

		// duplicates in the wordlist
		{[]string{"kite", "kite", "kite"}, []string{"kite", "kote"}, []string{"kite", "kite"}},

		// no matches at all
		{[]string{"abc"}, []string{"xyz", "ab", "abcd"}, []string{"", "", ""}},
		// repeated queries stay independent
		{[]string{"Hare", "hare"}, []string{"hare", "Hare", "HARE", "hore", "hare"}, []string{"hare", "Hare", "Hare", "Hare", "hare"}},
	}

	for _, tt := range tests {
		got := spellchecker(slices.Clone(tt.wordlist), slices.Clone(tt.queries))
		if !slices.Equal(got, tt.want) {
			t.Errorf("spellchecker(%q, %q) = %q, want %q", tt.wordlist, tt.queries, got, tt.want)
		}
	}
}
