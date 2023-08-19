package main

// https://leetcode.com/problems/unique-paths/
// Runtime 0 ms Beats 100% Memory 2.2 MB Beats 21.7%
func uniquePaths(m int, n int) int {
	visited := make([][]int, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]int, n)
	}

	return findPath(0, 0, visited)
}

func findPath(row, col int, visited [][]int) int {
	if row == len(visited)-1 && col == len(visited[0])-1 {
		return 1
	}

	if visited[row][col] > 0 {
		return visited[row][col]
	}

	var count int
	if row < len(visited)-1 {
		count += findPath(row+1, col, visited)
	}

	if col < len(visited[0])-1 {
		count += findPath(row, col+1, visited)
	}

	visited[row][col] = count

	return count
}
