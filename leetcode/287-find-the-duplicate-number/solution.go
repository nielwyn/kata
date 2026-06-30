// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
//
// There is only one repeated number in nums, return this repeated number.
//
// You must solve the problem without modifying the array nums and using only constant extra space.
// package findtheduplicatenumber
//
// Input: nums = [1,3,4,2,2]
// Output: 2
// Input: nums = [3,1,3,4,2]
// Output: 3
// Input: nums = [3,3,3,3,3]
// Output: 3
//
// 1 <= n <= 105
// nums.length == n + 1
// 1 <= nums[i] <= n
// All the integers in nums appear only once except for precisely one integer which appears two or more times.

package findtheduplicatenumber

// O(n log n) time, O(1) space
func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// TODO: two pointer
