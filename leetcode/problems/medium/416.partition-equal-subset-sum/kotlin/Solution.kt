// https://leetcode.com/problems/partition-equal-subset-sum/description/
class Solution {
    fun canPartition(nums: IntArray): Boolean {
        val total = nums.sum()
        if (total % 2 != 0) return false
        val target = total / 2

        if (nums.any { it > target }) return false

        nums.sortDescending()

        val dp = BooleanArray(target + 1)
        dp[0] = true

        var reachable = 0

        for (num in nums) {
            val newReachable = minOf(reachable + num, target)
            for (i in newReachable downTo num) {
                dp[i] = dp[i] || dp[i - num]
            }
            reachable = newReachable

            if (dp[target]) return true
        }

        return false
    }
}