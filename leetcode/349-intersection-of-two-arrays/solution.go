// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.
//
// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]
//
// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.
//
// Constraints:
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000

package intersectionoftwoarrays

// O(n+m) time, O(n) space
// func intersection(nums1 []int, nums2 []int) []int {
// 	inNums1 := make(map[int]bool, len(nums1))
// 	added := make(map[int]bool, len(nums2))
// 	result := []int{}
// 	for _, n := range nums1 {
// 		inNums1[n] = true
// 	}
// 	for _, n := range nums2 {
// 		if added[n] {
// 			continue
// 		}
// 		added[n] = true
// 		if inNums1[n] {
// 			result = append(result, n)
// 		}
// 	}
// 	return result
// }

// O(n+m) time, O(1) space
func intersection(nums1 []int, nums2 []int) []int {
	var inNums1 [1001]bool
	result := []int{}
	for _, n := range nums1 {
		inNums1[n] = true
	}
	for _, n := range nums2 {
		if inNums1[n] {
			result = append(result, n)
			inNums1[n] = false
		}
	}
	return result
}
