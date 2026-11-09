// Given a string s, find the length of the longest substring without duplicate
// characters.
//
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and
// "cab" are also correct answers.
//
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a
// substring.
//
// Constraints:
// 0 <= s.length <= 10⁵
// s consists of English letters, digits, symbols and spaces.

package longestsubstringwithoutrepeatingcharacters

// O(n) time, O(1) space
func lengthOfLongestSubstring(s string) int {
	var start, longest int
	visited := make(map[rune]int)
	for i, r := range s {
		if v, ok := visited[r]; ok && v >= start {
			start = v + 1
		}
		visited[r] = i
		longest = max(longest, i-start+1)
	}
	return longest
}
