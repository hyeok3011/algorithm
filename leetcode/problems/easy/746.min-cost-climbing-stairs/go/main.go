package main

// https://leetcode.com/problems/min-cost-climbing-stairs/description/
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
    dp := make([]int, len(cost) + 1)
	dp[0] = 0
	dp[1] = cost[0]
	for i := 1; i < len(cost); i++ {
        j := i+1
        dp[j] = min(dp[j-1] + cost[i], dp[j-2] + cost[i])
	}
	
	return dp[len(dp) - 1]
}
