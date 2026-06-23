// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
//
// Return any array that satisfies this condition.
//
// Input: nums = [3,1,2,4]
// Output: [2,4,3,1]
// Input: nums = [0]
// Output: [0]

// Constraints:
// 1 <= nums.length <= 5000
// 0 <= nums[i] <= 5000

package sortarraybyparity

// O(n) time, 0(1) space
func sortArrayByParity(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]%2 != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			right--
			continue
		}
		left++
	}
	return nums
}
