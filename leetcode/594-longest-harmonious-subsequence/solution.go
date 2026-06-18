// We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
//
// Given an integer array nums, return the length of its longest harmonious among all its possible subsequences.
// package longestharmonioussubsequence
//
// Input: nums = [1,3,2,2,5,2,3,7]
// Output: 5
// Input: nums = [1,2,3,4]
// Output: 2
// Input: nums = [1,1,1,1]
// Output: 0

package longestharmonioussubsequence

import (
	"sort"
)

func findLHS(nums []int) int {
	sort.Ints(nums)
	maxLen := 0
	l, r := 0, 1
	for r < len(nums) {
		diff := nums[r] - nums[l]
		if diff == 1 {
			maxLen = max(maxLen, r-l+1)
		}
		if diff <= 1 {
			r++
		} else {
			l++
		}
	}
	return maxLen
}
