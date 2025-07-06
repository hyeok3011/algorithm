// https://leetcode.com/problems/minimum-size-subarray-sum/description/

class Solution {
    fun minSubArrayLen(target: Int, nums: IntArray): Int {
        var minSize = Int.MAX_VALUE
        var left = 0
        var sum = 0L

        for (right in nums.indices) {
            sum += nums[right]

            while (sum >= target) {
                minSize = min(minSize, right - left + 1)
                sum -= nums[left]
                left += 1
            }
        }

        return if (minSize == Int.MAX_VALUE) 0 else minSize
    }
}