// Given an array of strings words, return the words that can be typed using
// letters of the alphabet on only one row of American keyboard like the image below.
//
// Note that the strings are case-insensitive, both lowercased and uppercased
// of the same letter are treated as if they are at the same row.
//
// In the American keyboard:
// the first row consists of the characters "qwertyuiop",
// the second row consists of the characters "asdfghjkl", and
// the third row consists of the characters "zxcvbnm".
//
// Input: words = ["Hello","Alaska","Dad","Peace"]
// Output: ["Alaska","Dad"]
// Both "a" and "A" are in the 2nd row of the American keyboard due to case insensitivity.

package keyboardrow

import "strings"

var top = map[rune]bool{
	'q': true, 'w': true, 'e': true, 'r': true, 't': true,
	'y': true, 'u': true, 'i': true, 'o': true, 'p': true,
}
var middle = map[rune]bool{
	'a': true, 's': true, 'd': true, 'f': true, 'g': true,
	'h': true, 'j': true, 'k': true, 'l': true,
}
var bottom = map[rune]bool{
	'z': true, 'x': true, 'c': true, 'v': true,
	'b': true, 'n': true, 'm': true,
}

func findWords(words []string) []string {
	var output []string
	for _, w := range words {
		var checkrow map[rune]bool
		valid := true
		for _, r := range strings.ToLower(w) {
			if checkrow == nil {
				switch {
				case top[r]:
					checkrow = top
					continue
				case middle[r]:
					checkrow = middle
					continue
				case bottom[r]:
					checkrow = bottom
					continue
				}
			}
			if _, ok := checkrow[r]; !ok {
				valid = false
				break
			}
		}
		if valid {
			output = append(output, w)
		}
	}

	return output
}
