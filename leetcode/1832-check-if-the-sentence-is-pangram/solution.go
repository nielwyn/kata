// A pangram is a sentence where every letter of the English alphabet appears at least once.
//
// Given a string sentence containing only lowercase English letters, return true
// if sentence is a pangram, or false otherwise.
//
// Example 1:
// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.
// Example 2:
// Input: sentence = "leetcode"
// Output: false
//
// Constraints:
// 1 <= sentence.length <= 1000
// sentence consists of lowercase English letters.

package checkifthesentenceispangram

// O(n) time, O(1) space
func checkIfPangram(sentence string) bool {
	visited := make(map[rune]struct{}, 26)
	for _, v := range sentence {
		visited[v] = struct{}{}
	}
	return len(visited) == 26
}

// O(n) time, O(1) space
// func checkIfPangram(sentence string) bool {
// 	seen := 0
// 	for i := range sentence {
// 		seen |= 1 << (sentence[i] - 'a')
// 		if seen == 1<<26-1 {
// 			return true
// 		}
// 	}
// 	return false
// }
