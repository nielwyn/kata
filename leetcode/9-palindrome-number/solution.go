// Given an integer x, return true if x is a , and false otherwise.
// An integer is a palindrome when it reads the same forward and backward.

package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}
