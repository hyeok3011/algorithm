// https://leetcode.com/problems/coin-change/description/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun coinChange(coins: IntArray, amount: Int): Int {
        val dp = IntArray(amount + 1)
        for (i in 0 until dp.size) {
            dp[i] = amount + 1
        }
        dp[0] = 0

        for(i in 0 until coins.size) {
            val coin = coins[i]
            for(j in coin until dp.size) {
                dp[j] = min(dp[j - coin] + 1, dp[j])
            }
        }
        if (dp[amount] == amount + 1) {
            return -1
        }
        return dp[amount]
    } 
}
