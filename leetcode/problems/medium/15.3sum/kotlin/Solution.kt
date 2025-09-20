// https://leetcode.com/problems/3sum/

class Solution {
    fun threeSum(nums: IntArray): List<List<Int>> {
        nums.sort()
        val answer = mutableListOf<List<Int>>()
        for (i in 0 until nums.size - 2) {
            if (nums[i] > 0) break
            if (i != 0 && nums[i - 1] == nums[i]) continue
            
            var left = i + 1
            var right = nums.size - 1
            while (left < right) {
                val sum = nums[i] + nums[left] + nums[right]
                when {
                    sum < 0 -> left++
                    sum > 0 -> right--
                    else -> {
                        answer.add(listOf(nums[i], nums[left], nums[right]))
                        val leftVal = nums[left]
                        val rightVal = nums[right]
                        while (left < right && nums[left] == leftVal) left++
                        while (left < right && nums[right] == rightVal) right--
                    }
                }
            }
        }
        return answer
    }
}
