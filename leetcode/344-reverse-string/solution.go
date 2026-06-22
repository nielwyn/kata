// Write a function that reverses a string. The input string is given as an array of characters s.
//
// You must do this by modifying the input array in-place with O(1) extra memory.
//
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

// Constraints:
// 1 <= s.length <= 105
// s[i] is a printable ascii character.

package reversestring

// O(n) time, O(1) space
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
