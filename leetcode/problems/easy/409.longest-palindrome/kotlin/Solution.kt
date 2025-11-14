// https://leetcode.com/problems/longest-palindrome/description/

class Solution {
    fun longestPalindrome(s: String): Int {
        val letterCounts = IntArray(128)
        for (char in s) {
            letterCounts[char.code]++
        }

        val evenParts = letterCounts.sumOf { it / 2 * 2 }
        val hasOdd = letterCounts.any { it % 2 == 1 }
        return if (hasOdd) evenParts + 1 else evenParts
    }
}