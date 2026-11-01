// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
//
// Example 1:
// Input: nums = [3,0,1]
// Output: 2
//
// Example 2:
// Input: nums = [0,1]
// Output: 2
//
// Example 3:
// Input: nums = [9,6,4,2,3,5,7,0,1]
// Output: 8
//
// Constraints:
// n == nums.length
// 1 <= n <= 10⁴
// 0 <= nums[i] <= n
// All the numbers of nums are unique.

package missingnumber

// O(n log n) time, O(log n) space
// func missingNumber(nums []int) int {
// 	slices.Sort(nums)
// 	n := 0
// 	for _, v := range nums {
// 		if n != v {
// 			return n
// 		}
// 		n++
// 	}
// 	return n
// }
//
//

// O(n) time, O(1) space
func missingNumber(nums []int) int {
	n := len(nums)
	want := n * (n + 1) / 2
	for _, v := range nums {
		want -= v
	}
	return want
}
