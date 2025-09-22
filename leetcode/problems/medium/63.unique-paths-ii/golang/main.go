package main

// https://leetcode.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp[0][0] = 1
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row > 0 && obstacleGrid[row][col] == 0 {
				dp[row][col] += dp[row-1][col]
			}

			if col > 0 && obstacleGrid[row][col] == 0 {
				dp[row][col] += dp[row][col-1]
			}
		}
	}
	return dp[m-1][n-1]
}
