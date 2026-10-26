// Given an integer numRows, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly
// above it.
//
// Example 1:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
// Example 2:
// Input: numRows = 1
// Output: [[1]]
//
// Constraints:
// 1 <= numRows <= 30

package pascalstriangle

// O(n²) time, O(n) auxiliary space
// func generate(numRows int) [][]int {
// 	triangle := [][]int{{1}}
// 	for i := range numRows - 1 {
// 		padded := slices.Concat([]int{0}, triangle[i], []int{0})
// 		row := []int{}
// 		for j := range len(padded) - 1 {
// 			row = append(row, padded[j]+padded[j+1])
// 		}
// 		triangle = append(triangle, row)
// 	}
// 	return triangle
// }

// O(n²) time, O(1) auxiliary space
func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range numRows {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle[i] = row
	}
	return triangle
}
