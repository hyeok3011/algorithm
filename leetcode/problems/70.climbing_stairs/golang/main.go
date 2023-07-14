package main

// https://leetcode.com/problems/climbing-stairs/
// Runtime 0 ms Beats 100% Memory 1.9 MB Beats 30.46%
// 풀어보니 dp문제
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
