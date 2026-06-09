// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	prefixLength := 1
	prefix := ""

	var check func()
	check = func() {
		for i, v := range strs {
			if len(v) < prefixLength {
				return
			}
			if strs[0][:prefixLength] != v[:prefixLength] {
				return
			}
			if i == len(strs)-1 {
				prefix = strs[0][:prefixLength]
			}

		}
		prefixLength += 1
		check()
	}
	check()

	return prefix
}
