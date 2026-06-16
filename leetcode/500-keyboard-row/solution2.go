package keyboardrow

import "strings"

var rowOf [26]int

func init() {
	for _, r := range "qwertyuiop" {
		rowOf[r-'a'] = 1
	}
	for _, r := range "asdfghjkl" {
		rowOf[r-'a'] = 2
	}
	for _, r := range "zxcvbnm" {
		rowOf[r-'a'] = 3
	}
}

func findWords2(words []string) []string {
	var output []string
	for _, w := range words {
		if singleRow(w) {
			output = append(output, w)
		}
	}
	return output
}

func singleRow(w string) bool {
	lower := strings.ToLower(w)
	row := rowOf[lower[0]-'a']
	for _, r := range lower[1:] {
		if rowOf[r-'a'] != row {
			return false
		}
	}
	return true
}
