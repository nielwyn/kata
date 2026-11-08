// In an alien language, surprisingly, they also use English lowercase letters,
// but possibly in a different order. The order of the alphabet is some permutation
// of lowercase letters.
//
// Given a sequence of words written in the alien language, and the order of the
// alphabet, return true if and only if the given words are sorted
// lexicographically in this alien language.
//
// Example 1:
// Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
// Output: true
// Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
//
// Example 2:
// Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
// Output: false
// Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1],
// hence the sequence is unsorted.
//
// Example 3:
// Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
// Output: false
// Explanation: The first three characters "app" match, and the second string
// is shorter (in size.) According to lexicographical rules "apple" > "app",
// because 'l' > '∅', where '∅' is defined as the blank character which is less
// than any other character.
//
// Constraints:
// 1 <= words.length <= 100
// 1 <= words[i].length <= 20
// order.length == 26
// All characters in words[i] and order are English lowercase letters.

package verifyinganaliendictionary

// O(n*k) time, O(1) space
func isAlienSorted(words []string, order string) bool {
	rank := make(map[rune]int, 26)
	for i, letter := range order {
		rank[letter] = i
	}
	for i := 1; i < len(words); i++ {
		if outOfOrder(words[i-1], words[i], rank) {
			return false
		}
	}
	return true
}

func outOfOrder(prev, next string, rank map[rune]int) bool {
	for i := 0; i < len(prev) && i < len(next); i++ {
		if prev[i] != next[i] {
			return rank[rune(prev[i])] > rank[rune(next[i])]
		}
	}
	return len(prev) > len(next)
}
