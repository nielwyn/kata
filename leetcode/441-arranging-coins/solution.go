// You have n coins and you want to build a staircase with these coins. The
// staircase consists of k rows where the ith row has exactly i coins. The last
// row of the staircase may be incomplete.
//
// Given the integer n, return the number of complete rows of the staircase you will build.
//
// Input: n = 5
// Output: 2
// Explanation: Because the 3rd row is incomplete, we return 2.
// Input: n = 8
// Output: 3
// Explanation: Because the 4th row is incomplete, we return 3.
//
// Constraints:
// 1 <= n <= 231 - 1

package arrangingcoins

// O(log n) time, O(1) space
func arrangeCoins(n int) int {
	left, right := 1, n
	k := 0
	for left <= right {
		mid := (left + right) / 2
		if mid*(mid+1)/2 <= n {
			k = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return k
}

// O(sqrt(n)) time, O(1) space
// func arrangeCoins(n int) int {
// 	k := 1
// 	for n > k*(k+1)/2 {
// 		k++
// 	}
// 	if n == k*(k+1)/2 {
// 		return k
// 	}
// 	return k - 1
// }
