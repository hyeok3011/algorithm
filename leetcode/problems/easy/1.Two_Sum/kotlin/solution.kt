class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val memo = mutableMapOf<Int, Int>()
        nums.forEachIndexed { i, v ->
            val complement = target - v
            memo[complement]?.let { complementIndex ->
                return intArrayOf(complementIndex, i)
            }
            memo[v] = i
        }
        throw IllegalArgumentException("No two sum solution")
    }
}
