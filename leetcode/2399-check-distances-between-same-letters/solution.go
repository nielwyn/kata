// You are given a 0-indexed string s consisting of only lowercase English
// letters, where each letter in s appears exactly twice. You are also given a
// 0-indexed integer array distance of length 26.
//
// Each letter in the alphabet is numbered from 0 to 25
// (i.e. 'a' -> 0, 'b' -> 1, 'c' -> 2, ... , 'z' -> 25).
//
// In a well-spaced string, the number of letters between the two occurrences
// of the ith letter is distance[i]. If the ith letter does not appear in s,
// then distance[i] can be ignored.
//
// Return true if s is a well-spaced string, otherwise return false.
//
// Input: s = "abaccb", distance = [1,3,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
// Output: true

package checkdistancesbetweensameetters

func checkDistances(s string, distance []int) bool {
	visited := make(map[rune]int)
	for i, r := range s {
		if _, ok := visited[r]; ok {
			visited[r] = i - visited[r] - 1
			continue
		}
		visited[r] = i
	}
	for k, v := range visited {
		if v != distance[k-'a'] {
			return false
		}
	}
	return true
}

// func checkDistances(s string, distance []int) bool {
// 	var first [26]int
// 	var seen [26]bool
//
// 	for i, r := range s {
// 		idx := r - 'a'
// 		if seen[idx] {
// 			if i-first[idx]-1 != distance[idx] {
// 				return false
// 			}
// 			continue
// 		}
// 		first[idx] = i
// 		seen[idx] = true
// 	}
// 	return true
// }
