package main

/*
https://leetcode.com/problems/edit-distance/
최근에 푼 알고리즘 문제중 제일 어려웠다.... 역시 DP문제가 제일 어렵다
먼저 작은 문자들 끼리 비교하며 올라가서 최종 결과를 봐야하기 때문에 BottomUp방식을 사용했다.
풀이를 깔끔하게 정리해놓은 사람이 있다.
https://leetcode.com/problems/edit-distance/solutions/3231533/golang-dynamic-programming-with-explanation/
나중에 다시 풀어볼만한 문제다.
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

func min(data ...int) int {
	minVal := data[0]
	for _, v := range data {
		if minVal > v {
			minVal = v
		}
	}

	return minVal
}
