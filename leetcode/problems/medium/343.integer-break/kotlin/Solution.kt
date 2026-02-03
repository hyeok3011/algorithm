// https://leetcode.com/problems/integer-break/
// dp로 풀려다가 점화식 실패...
class Solution {
    // 2 = 1
    // 3 = 2   | 2 + 1
    // 4 = 4   | 2 + 2
    // 5 = 6   | 3 + 2
    // 6 = 9   | 3 + 3
    // 7 = 12  | 3 + 2 + 2
    // 8 = 18  | 3 + 3 + 2
    // 9 = 27  | 3 + 3 + 3
    // 10 = 36 | 3 + 3 + 2 + 2
    fun integerBreak(n: Int): Int {
        if (n == 2) {
            return 1
        }
        if (n == 3) {
            return 2
        }

        var product = 1
        var number = n
        while (number > 4) {
            product *= 3
            number -= 3
        }

        return product * number
    }

    // 다른사람의 DP방식 풀이
    fun integerBreakDP(n: Int): Int {
        val dp = IntArray(n+1)
        dp[1] = 1

        for (num in 2..n) {
            if (num != n) {
                dp[num] = num
            }
            for (i in 1 until num) {
                val prod = dp[i] * dp[num-i]
                dp[num] = maxOf(dp[num], prod)
            }
        }

        return dp[n]
    }
}
