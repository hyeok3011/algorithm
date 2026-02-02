// https://leetcode.com/problems/house-robber/
class Solution {
    fun rob(nums: IntArray): Int {
        if (nums.size == 1) {
            return nums[0]
        }

        var preMax = nums[0]
        var currentMax = nums[1]
        for(i in 2 until nums.size) {
            val newMoney = preMax + nums[i]
            preMax = max(currentMax, preMax)
            currentMax = max(currentMax, newMoney)
        }

        return max(currentMax, preMax)
    }
}