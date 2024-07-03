package main

/*
https://leetcode.com/problems/house-robber/description/
붙어있지 않은 엘리먼트의 합중에 가장 큰 수를 구하는 문제이다.
문제를 풀수있는 방법은
1. 모든 경우를 확인해 보는 것이다.
한단, 안한다, 안안한 의 모든 경우를 확인해 봐야한다.
2. DP형식으로 풀어보는것
sums[i] 기준으로 최선의 값을 선택하는 방법을 찾는것

총 2가지 방법으로 풀수 있을듯 하다.
먼저 1번 방법의 경우 너무 많은 경우의 수가 존재하고 nums길이는 총 400까지 될 수 있다.
그나마 2번 방법이 더 효율적이라고 판단하고 문제를 풀었다.

DP문제가 항상 제일 약한데 이번 문제는 그나마 풀만했다.
먼저 nums보다 1이나 2를 더 크게 dp를 만들어 준다. padding을 줘서 계산하기 쉽게 풀기 위함이다.
그 이후로는 nums[i]기준으로 최선의 판단을 하면된다.
[   2, 1, 1, 2]
[0, 0, 0, 0, 0]
으로 시작하여
[   2, 1, 1, 2]
[0, 2, 2, 3, 4]
*/
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
