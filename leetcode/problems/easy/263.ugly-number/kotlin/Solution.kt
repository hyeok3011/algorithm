// https://leetcode.com/problems/ugly-number/
class Solution {
    fun isUgly(n: Int): Boolean {
        if (n <= 0) return false

        var number = n
        val factors = intArrayOf(2, 3, 5)
        for (factor in factors) {
            number = removeFactor(number, factor)
        }

        return number == 1
    }

    fun removeFactor(number: Int, factor: Int): Int {
        var current = number
        while(current % factor == 0) {
            current /= factor
        }

        return current
    }
}
