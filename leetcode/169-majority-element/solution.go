// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
//
// Constraints:
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
// The input is generated such that a majority element will exist in the array.
//
// Follow-up: Could you solve the problem in linear time and in O(1) space?

package majorityelement

// O(n) time, O(n) space
// func majorityElement(nums []int) int {
// 	visited := map[int]int{}
// 	for _, v := range nums {
// 		visited[v]++
// 		if visited[v] > len(nums)/2 {
// 			return v
// 		}
// 	}
// 	panic("unreachable")
// }

// Boyer–Moore Majority Vote Algorithm (majority element is guaranteed)
// O(n) time, O(1) space
// func majorityElement(nums []int) int {
// 	var candidate, count int
// 	// pass 1 only
// 	for _, num := range nums {
// 		if count == 0 {
// 			candidate = num
// 		}
// 		if num == candidate {
// 			count++
// 		} else {
// 			count--
// 		}
// 	}
// 	return candidate
// }

// Boyer–Moore Majority Vote Algorithm (majority element not guaranteed)
// O(n) time, O(1) space
func majorityElement(nums []int) int {
	var candidate, count int
	// pass 1: find the only possible candidate
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// pass 2: verify candidate appears > n/2 times
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count <= len(nums)/2 {
		panic("no majority element")
	}
	return candidate
}
