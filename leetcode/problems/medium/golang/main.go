package main

import "math"

// https://leetcode.com/problems/minimum-path-sum/description/
// Runtime 4 ms Beats 95.61% Memory 4.5 MB Beats 17.24%
func minPathSum(grid [][]int) int {
	used := make([][]int, len(grid))
	for i := 0; i < len(used); i++ {
		used[i] = make([]int, len(grid[0]))
		for j := 0; j < len(used[i]); j++ {
			used[i][j] = -1
		}
	}
	return backTracking(0, 0, grid, used)
}

func backTracking(row, col int, grid [][]int, used [][]int) int {
	if row == len(grid)-1 && col == len(grid[0])-1 {
		return grid[row][col]
	}
	if used[row][col] != -1 {
		return used[row][col]
	}

	minPath := math.MaxInt32

	if row+1 < len(grid) {
		minPath = min(minPath, backTracking(row+1, col, grid, used))
	}

	if col+1 < len(grid[0]) {
		minPath = min(minPath, backTracking(row, col+1, grid, used))
	}

	used[row][col] = minPath + grid[row][col]
	return used[row][col]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// DP로 문제를 풀었다 이렇게 풀 생각을 못했는데...
func AnotherSolution(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
