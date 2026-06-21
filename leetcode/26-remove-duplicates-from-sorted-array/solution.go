// Given an integer array nums sorted in non-decreasing order, remove the
// duplicates in-place such that each unique element appears only once. The
// relative order of the elements should be kept the same.
//
// Consider the number of unique elements in nums to be k. After removing
// duplicates, return the number of unique elements k.
//
// The first k elements of nums should contain the unique numbers in sorted order.
// The remaining elements beyond index k - 1 can be ignored.
//
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
// Input: nums = [0,0,1,1,1,2,2,3,3,4]
// Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

package removeduplicatesfromsortedarray

// O(n) time, O(1) space
func removeDuplicates(nums []int) int {
	k := 1
	for i := range nums {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
		continue
	}
	return k
}
