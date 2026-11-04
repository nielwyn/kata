// Given an integer array nums of length n where all the integers of nums are
// in the range [1, n] and each integer appears at most twice, return an array
// of all the integers that appears twice.
//
// You must write an algorithm that runs in O(n) time and uses only constant
// auxiliary space, excluding the space needed to store the output
//
// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [2,3]
//
// Example 2:
// Input: nums = [1,1,2]
// Output: [1]
//
// Example 3:
// Input: nums = [1]
// Output: []
//
// Constraints:
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
// Each element in nums appears once or twice.

package findallduplicatesinanarray

// O(n) time, O(1) space
func findDuplicates(nums []int) []int {
	want := []int{}
	for i := range nums {
		v := nums[i]
		// negate inline instead of match.Abs
		if v < 0 {
			v = -v
		}
		if nums[v-1] < 0 {
			want = append(want, v)
			continue
		}
		nums[v-1] = -nums[v-1]
	}
	return want
}
