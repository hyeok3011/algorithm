package main

// https://leetcode.com/problems/domino-and-tromino-tiling/?envType=study-plan-v2&envId=leetcode-75
// DP로 풀긴했으나 사람들의 힌트를 보고 풀었음.
// https://leetcode.com/problems/domino-and-tromino-tiling/solutions/4581905/gopython-dp-approach-on-o1-space-with-ex-kbrc/?envType=study-plan-v2&envId=leetcode-75
// @@@@@@@@@@@@@@@@

func numTilings(n int) int {
    if n <= 2 {
        return n
    }

    dp := make([]int, n + 1)
    dp[0] = 1
    dp[1] = 1
    dp[2] = 2
    for i := 3; i <= n; i++ {
        dp[i] = (2 * dp[i-1] + dp[i-3]) % 1000000007
    }
    return dp[n]
}
