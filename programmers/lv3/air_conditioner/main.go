package main

// https://school.programmers.co.kr/learn/courses/30/lessons/214289#qna
// greedy방법은 테스트 케이스 3개가 실패했다 ㅠ
// 완전 탐색은 3 ** 1000가 될거같아 포기했다.
// 그래서 결국 DP방법으로 풀어봤다.
// 최대한 간소화 하기위해 maxDiff를 계산하여 적정온도에서 -1까지만 계산하여 풀었다.

func solution(outTemperature, lowTemperature, highTemperature, diffTempPower, sameTempPower int, onboard []int) int {
	maxDiff := max(lowTemperature-outTemperature, outTemperature-highTemperature) + 2
	dp := make([][]int, len(onboard)+1) // add padding
	inf := 2<<31 - 1
	for i := range dp {
		temperatureLine := make([]int, maxDiff)
		for j := range temperatureLine {
			temperatureLine[j] = inf
		}
		dp[i] = temperatureLine
	}

	dp[0][maxDiff-1] = 0

	for i := 0; i < len(onboard); i++ {
		start, end := 0, maxDiff
		if onboard[i] == 1 {
			start, end = 0, 2
		}
		for temp := start; temp < end; temp++ {
			if dp[i][temp] == inf {
				continue
			}
			currentPower := dp[i][temp]

			// on same temperature
			dp[i+1][temp] = min(dp[i+1][temp], currentPower+sameTempPower)

			// on diffrent temperature
			if temp != 0 {
				dp[i+1][temp-1] = min(dp[i+1][temp-1], currentPower+diffTempPower)
			}

			// off
			if temp == maxDiff-1 {
				dp[i+1][temp] = min(dp[i+1][temp], currentPower)
			} else {
				dp[i+1][temp+1] = min(dp[i+1][temp+1], currentPower)
			}
		}
	}

	minVal := inf
	for _, v := range dp[len(onboard)] {
		minVal = min(minVal, v)
	}

	return minVal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
