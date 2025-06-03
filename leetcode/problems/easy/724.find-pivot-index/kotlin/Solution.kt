
// https://leetcode.com/problems/find-pivot-index/description/?envType=study-plan-v2&envId=leetcode-75
class Solution {
    fun pivotIndex(nums: IntArray): Int {
        var totalSum = nums.sum()
        var preSum = 0
        for ((i, num) in nums.withIndex()) {
            totalSum -= num
            if (totalSum == preSum) {
                return i
            }
            preSum += num
        }
        return -1
    }
}