// Given an integer array nums and an integer k, return true if there are two
// distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
//
// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: true
//
// Example 2:
// Input: nums = [1,0,1,1], k = 1
// Output: true
//
// Example 3:
// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false
//
// Constraints:
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 0 <= k <= 10⁵

package containsduplicate2

// O(n) time, O(n) space
func containsNearbyDuplicate(nums []int, k int) bool {
	visited := map[int]int{}
	for i, v := range nums {
		if j, ok := visited[v]; ok && i-j <= k {
			return true
		}
		visited[v] = i
	}
	return false
}
