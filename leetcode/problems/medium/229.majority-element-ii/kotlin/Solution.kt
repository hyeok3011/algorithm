// https://leetcode.com/problems/majority-element-ii/description/
// 다른사람의 majority vote 방식을 보고 다시 풀이
class Solution {
    fun majorityElement(nums: IntArray): List<Int> {
        var candidateA = 0
        var candidateB = 0
        var countA = 0
        var countB = 0

        for (num in nums) {
            when {
                candidateA == num -> countA++
                candidateB == num -> countB++
                countA == 0 -> { candidateA = num; countA++ }
                countB == 0 -> { candidateB = num; countB++ }
                else -> { countA--; countB-- }
            }
        }

        countA = 0
        countB = 0
        for (num in nums) {
            if (candidateA == num) countA++
            else if (candidateB == num) countB++
        }

        val answer = mutableListOf<Int>()
        if (countA > nums.size / 3) answer.add(candidateA)
        if (countB > nums.size / 3) answer.add(candidateB)
        return answer
    }
}
