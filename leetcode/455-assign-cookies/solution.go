// Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie.
//
// Each child i has a greed factor g[i], which is the minimum size of a cookie that the child will be content with; and each cookie j has a size s[j]. If s[j] >= g[i], we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.
//
// Input: g = [1,2,3], s = [1,1]
// Output: 1
// Input: g = [1,2], s = [3,2,1]
// Output: 2
//
// Constraints:
// 1 <= g.length <= 3 * 104
// 0 <= s.length <= 3 * 104
// 1 <= g[i], s[j] <= 231 - 1

package assigncookies

import (
	"slices"
)

// O(n log n + m log m) time, O(log n + log m) space
func findContentChildren(g []int, s []int) int {
	gp, sp, content := 0, 0, 0
	slices.Sort(g)
	slices.Sort(s)
	for gp < len(g) && sp < len(s) {
		if s[sp] >= g[gp] {
			content++
			gp++
			sp++
			continue
		}
		sp++
	}
	return content
}
