// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true
//
// Example 2:
// Input: s = "rat", t = "car"
// Output: false
//
// Constraints:
// 1 <= s.length, t.length <= 5 * 10⁴
// s and t consist of lowercase English letters.
//
// Follow up: What if the inputs contain Unicode characters? How would you adapt
// your solution to such a case?

package validanagram

// lowercase English only
// O(n) time, O(1) space
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	var count [26]int
// 	for i := range s {
// 		count[s[i]-'a']++
// 		count[t[i]-'a']--
// 	}
// 	for _, c := range count {
// 		if c != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// Unicode (follow up)
// O(n) time, O(k) space for k distinct runes
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	visited := map[rune]int{}
	for _, r := range s {
		visited[r]++
	}
	for _, r := range t {
		if visited[r] == 0 {
			return false
		}
		visited[r]--
	}
	return true
}
