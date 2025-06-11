// https://leetcode.com/problems/max-consecutive-ones-iii/?envType=study-plan-v2&envId=leetcode-75
class Solution {
    fun longestOnes(nums: IntArray, k: Int): Int {
        var left = 0
        var zeroCount = 0
        var maxLength = 0
        
        for (i in nums) {
            if (nums[i] == 0) {
                zeroCount++
            }
            
            while (zeroCount > k) {
                if (nums[left] == 0) {
                    zeroCount--
                }
                left++
            }
            
            maxLength = maxOf(maxLength, i - left + 1)
        }
        
        return maxLength
    }
}

// 다른 사람의 풀이법. 똑똑스..
class Solution {
    fun longestOnes(nums: IntArray, k: Int): Int {
        var left = 0
        var k = k

        for (right in nums.indices) {
            if (nums[right] == 0) {
                k--
            }

            if (k < 0) {
                if (nums[left] == 0) {
                    k++
                }
                left++
            }
        }
        return nums.size - left
    }
}