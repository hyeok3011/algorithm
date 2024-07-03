package main

// https://leetcode.com/problems/pascals-triangle/description/
// Runtime 0 ms Beats 100% Memory 2.3 MB Beats 39.79%
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	triangle := make([][]int, numRows)

	triangle[0] = []int{1}
	triangle[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		current := make([]int, i+1)
		current[0] = 1
		current[i] = 1

		for j := 1; j < i; j++ {
			current[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		triangle[i] = current
	}

	return triangle
}
