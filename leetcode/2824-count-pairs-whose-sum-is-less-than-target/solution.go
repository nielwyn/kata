// Given a 0-indexed integer array nums of length n and an integer target,
// return the number of pairs (i, j) where 0 <= i < j < n and nums[i] + nums[j] < target.
//
// Example 1:
// Input: nums = [-1,1,2,3,1], target = 2
// Output: 3
// Example 2:
// Input: nums = [-6,2,5,-2,-7,-1,3], target = -2
// Output: 10
//
// Constraints:
// 1 <= nums.length == n <= 50
// -50 <= nums[i], target <= 50

package countpairswhosesumislessthantarget

import "slices"

// Binary Search
// O(n log n) time, O(log n) space
// func countPairs(nums []int, target int) int {
// 	pairs := 0
// 	slices.Sort(nums)
// 	for i, v := range nums {
// 		left, right := i+1, len(nums)-1
// 		for left <= right {
// 			mid := (left + right) / 2
// 			if v+nums[mid] < target {
// 				pairs += mid - left + 1
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		}
// 	}
// 	return pairs
// }

// Two Pointers
// O(n log n) time, O(log n) space
func countPairs(nums []int, target int) int {
	pairs := 0
	left, right := 0, len(nums)-1
	slices.Sort(nums)
	for left < right {
		if nums[left]+nums[right] < target {
			pairs += right - left
			left++
		} else {
			right--
		}
	}
	return pairs
}
