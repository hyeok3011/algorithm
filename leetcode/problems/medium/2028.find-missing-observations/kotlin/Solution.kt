// https://leetcode.com/problems/find-missing-observations/
class Solution {
    fun missingRolls(rolls: IntArray, mean: Int, n: Int): IntArray {
        val knownSum = rolls.sum()
        val totalSum = (rolls.size + n) * mean
        val diff = totalSum - knownSum

        if (diff > 6 * n || diff < n) {
            return intArrayOf()
        }

        val avg = diff / n
        var remainder = diff % n
        val missingValues = IntArray(n)
        for (i in 0 until n) {
            missingValues[i] = avg
            if (remainder > 0) {
                missingValues[i]++
                remainder--
            }
        }

        return missingValues
    }
}