// You are working in a ball factory where you have n balls numbered from lowLimit
// up to highLimit inclusive (i.e., n == highLimit - lowLimit + 1), and an
// infinite number of boxes numbered from 1 to infinity.
//
// Your job at this factory is to put each ball in the box with a number equal to
// the sum of digits of the ball's number. For example, the ball number 321 will
// be put in the box number 3 + 2 + 1 = 6 and the ball number 10 will be put in
// the box number 1 + 0 = 1.
//
// Given two integers lowLimit and highLimit, return the number of balls in the
// box with the most balls.
//
// Example 1:
// Input: lowLimit = 1, highLimit = 10
// Output: 2
//
// Example 2:
// Input: lowLimit = 5, highLimit = 15
// Output: 2
//
// Example 3:
// Input: lowLimit = 19, highLimit = 28
// Output: 2
//
// Constraints:
// 1 <= lowLimit <= highLimit <= 10⁵

package maximumnumberofballsinabox

// O(n) time, O(1) space
func countBalls(lowLimit int, highLimit int) int {
	var box [46]int // max digit sum: 99999 = 45
	maxVal := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := sumDigits(i)
		box[sum]++
		if box[sum] > maxVal {
			maxVal = box[sum]
		}
	}
	return maxVal
}

func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
