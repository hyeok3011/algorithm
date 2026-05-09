// # https://leetcode.com/problems/buddy-strings/description/
class Solution {
    fun buddyStrings(s: String, goal: String): Boolean {
        if (s.length != goal.length) return false

        if (s == goal) {
            return s.toSet().size < s.length
        }

        val diffs = s.indices.filter{s[it] != goal[it]}
        if (diffs.size != 2) return false

        val (i, j) = diffs
        return s[i] == goal[j] && s[j] == goal[i]
    }
}