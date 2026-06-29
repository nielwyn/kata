// Given a positive integer num, return true if num is a perfect square or false otherwise.
//
// A perfect square is an integer that is the square of an integer. In other
// words, it is the product of some integer with itself.
//
// You must not use any built-in library function, such as sqrt.
//
// Input: num = 16
// Output: true
// Input: num = 14
// Output: false
//
// Constraints:
// 1 <= num <= 2^31 - 1

package validperfectsquare

// O(log n) time, O(1) space
func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
