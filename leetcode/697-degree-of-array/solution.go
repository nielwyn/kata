// Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.
// Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.
// Input: nums = [1,2,2,3,1]
// Output: 2

package degreeofarray

func findShortestSubArray(nums []int) int {
	count := map[int]int{}
	first := map[int]int{}
	last := map[int]int{}

	for i, n := range nums {
		count[n]++
		if _, ok := first[n]; !ok {
			first[n] = i
		}
		last[n] = i
	}

	degree := 0
	shortestSubArr := 0
	for n, freq := range count {
		if freq > degree {
			degree = freq
			shortestSubArr = last[n] - first[n] + 1
		}
		if freq == degree {
			if last[n]-first[n] < shortestSubArr {
				shortestSubArr = last[n] - first[n] + 1
			}
		}
	}

	return shortestSubArr
}
