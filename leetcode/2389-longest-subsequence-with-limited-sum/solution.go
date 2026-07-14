// You are given an integer array nums of length n, and an integer array queries
// of length m.
//
// Return an array answer of length m where answer[i] is the maximum size of a
// subsequence that you can take from nums such that the sum of its elements is
// less than or equal to queries[i].
//
// A subsequence is an array that can be derived from another array by deleting
// some or no elements without changing the order of the remaining elements.
//
// Example 1:
// Input: nums = [4,5,2,1], queries = [3,10,21]
// Output: [2,3,4]
// Example 2:
// Input: nums = [2,3,4,5], queries = [1]
// Output: [0]
//
// Constraints:
// n == nums.length
// m == queries.length
// 1 <= n, m <= 1000
// 1 <= nums[i], queries[i] <= 10^6

package longestsubsequencewithlimitedsum

import "slices"

// O(n log n + m log n ) time, O(n) space
func answerQueries(nums []int, queries []int) []int {
	slices.Sort(nums)
	prefix := make([]int, len(nums))
	ans := make([]int, len(queries))
	sums := 0
	for i, v := range nums {
		sums += v
		prefix[i] = sums
	}
	for i, v := range queries {
		left, right := 0, len(prefix)-1
		for left <= right {
			mid := (left + right) / 2
			if prefix[mid] <= v {
				ans[i] = mid + 1
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return ans
}
