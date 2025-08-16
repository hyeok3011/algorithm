package main

func isInterleave(s1, s2, s3 string) bool {
	s1Len, s2Len, s3Len := len(s1), len(s2), len(s3)
	if s1Len+s2Len != s3Len {
		return false
	}

	dp := make([][]bool, s1Len+1)
	for i := range dp {
		dp[i] = make([]bool, s2Len+1)
	}

	dp[0][0] = true

	for i := 1; i <= s1Len; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= s2Len; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for row := 1; row <= s1Len; row++ {
		for col := 1; col <= s2Len; col++ {
			s1Match := dp[row-1][col] && s1[row-1] == s3[row+col-1]
			s2Match := dp[row][col-1] && s2[col-1] == s3[row+col-1]
			dp[row][col] = s1Match || s2Match
		}
	}

	return dp[s1Len][s2Len]
}
