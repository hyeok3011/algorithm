package main

// https://leetcode.com/problems/equal-row-and-column-pairs/?envType=study-plan-v2&envId=leetcode-75
func equalPairs(grid [][]int) int {
	rowSet := make(map[[200]int]int, len(grid))
	for row := range grid {
		numbers := [200]int{}
		copy(numbers[:], grid[row])
		rowSet[numbers]++
	}

	colSet := make(map[[200]int]int, len(grid[0]))
	for col := range grid[0] {
		numbers := [200]int{}
		for row := range grid {
			numbers[row] = grid[row][col]
		}
		colSet[numbers]++
	}

	answer := 0
	for key, _ := range rowSet {
		answer += rowSet[key] * colSet[key]
	}

	return answer
}
