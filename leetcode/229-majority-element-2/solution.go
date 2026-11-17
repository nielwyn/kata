// Given an integer array of size n, find all elements that appear more than ⌊n / 3⌋ times.
//
// Example 1:
// Input: nums = [3,2,3]
// Output: [3]
//
// Example 2:
// Input: nums = [1]
// Output: [1]
//
// Example 3:
// Input: nums = [1,2]
// Output: [1,2]
//
// Constraints:
// 1 <= nums.length <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹

package majorityelement2

// O(n) time, O(1) space
func majorityElement(nums []int) []int {
	const k = 3
	candidates := make(map[int]int, k-1)
	// pass 1
	for _, n := range nums {
		if _, ok := candidates[n]; ok {
			candidates[n]++
		} else if len(candidates) < k-1 {
			candidates[n] = 1
		} else {
			for val := range candidates {
				candidates[val]--
				if candidates[val] == 0 {
					delete(candidates, val)
				}
			}
		}
	}

	// pass 2
	for val := range candidates {
		candidates[val] = 0
	}
	for _, n := range nums {
		if _, ok := candidates[n]; ok {
			candidates[n]++
		}
	}
	want := make([]int, 0, k-1)
	for val, count := range candidates {
		if count > len(nums)/k {
			want = append(want, val)
		}
	}
	return want
}
