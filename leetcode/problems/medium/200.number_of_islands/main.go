package main

/*
https://leetcode.com/problems/number-of-islands/description/
grid에서 1은 육지를 뜻하고 0은 물을 뜻한다. 여기서 총 섬의 개수를 구하는 문제이다.
여기서 육지는 상하좌우로 연결된 육지를 하나의 섬으로 생각한다.
문제는 간단하다. grid를 순회하며 육지인경우 주변 같은 섬을 물로 바꿔주면 된다.
*/

func numIslands(grid [][]byte) int {
	islandCount := 0

	for i := range grid {
		for j := range grid[i] {
			if isIslandsDFS(grid, i, j) {
				islandCount++
			}
		}
	}

	return islandCount
}

func isIslandsDFS(grid [][]byte, row, col int) bool {
	if grid[row][col] == '0' {
		return false
	}
	grid[row][col] = '0'

	if row-1 >= 0 {
		isIslandsDFS(grid, row-1, col)
	}

	if row+1 < len(grid) {
		isIslandsDFS(grid, row+1, col)
	}

	if col-1 >= 0 {
		isIslandsDFS(grid, row, col-1)
	}

	if col+1 < len(grid[0]) {
		isIslandsDFS(grid, row, col+1)
	}

	return true
}
