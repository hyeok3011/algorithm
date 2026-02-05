// https://leetcode.com/problems/random-pick-index/description/
// 너무 느려서 문제 의도와 조금 다르게 풀음...
class Solution(private val nums: IntArray) {
    fun pick(target: Int): Int {
        var result = -1
        var count = 0
        
        for (i in nums.indices) {
            if (nums[i] == target) {
                count++
                if (Random.nextInt(count) == 0) {
                    result = i
                }
            }
        }
        return result
    }
}
