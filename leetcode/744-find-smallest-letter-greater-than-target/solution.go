// You are given an array of characters letters that is sorted in non-decreasing order, and a character target. There are at least two different characters in letters.
//
// Return the smallest character in letters that is lexicographically greater than target. If such a character does not exist, return the first character in letters.
//
// Input: letters = ["c","f","j"], target = "a"
// Output: "c"
// Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.
// Input: letters = ["c","f","j"], target = "c"
// Output: "f"
// Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.
// Input: letters = ["x","x","y","y"], target = "z"
// Output: "x"
// Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0].
//
// Constraints:
// 2 <= letters.length <= 10^4
// letters[i] is a lowercase English letter.
// letters is sorted in non-decreasing order.
// letters contains at least two different characters.
// target is a lowercase English letter.

package findsmallestlettergreaterthantarget

// O(log n) time, O(1) space
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right, ans := 0, len(letters)-1, letters[0]
	for left <= right {
		mid := (left + right) / 2
		if letters[mid] > target {
			ans = letters[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
