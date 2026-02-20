// https://leetcode.com/problems/4sum/
// 문제 자체는 3sum과 똑같은 문제지만 최적화하는 과정에서 조금 헷갈렸음.
class Solution {
    fun fourSum(nums: IntArray, target: Int): List<List<Int>> {
        nums.sort()
        val answer = mutableListOf<List<Int>>()
        val n = nums.size
        for (i in 0 until n - 3) {
            if (i != 0 && nums[i] == nums[i - 1]) continue
            if (nums[i].toLong() + nums[i + 1] + nums[i + 2] + nums[i + 3] > target) break
            if (nums[i].toLong() + nums[n - 3] + nums[n - 2] + nums[n - 1] < target) continue

            for (j in i+1 until n - 2) {
                if (j != i + 1 && nums[j] == nums[j-1]) continue
                if (nums[i].toLong() + nums[j] + nums[j + 1] + nums[j + 2] > target) break
                if (nums[i].toLong() + nums[j] + nums[n - 2] + nums[n - 1] < target) continue

                var left = j + 1
                var right = n - 1
                while (left < right) {
                    val sum = nums[i].toLong() + nums[j] + nums[left] + nums[right]
                    when {
                        sum < target -> left++
                        sum > target -> right--
                        else -> {
                            answer.add(listOf(nums[i], nums[j], nums[left], nums[right]))
                            val leftVal = nums[left]
                            while(left < right && leftVal == nums[left]) left++
                            val rightVal = nums[right]
                            while(left < right && rightVal == nums[right]) right--
                        }
                    }
                }
            }
        }
        return answer
    }
}
