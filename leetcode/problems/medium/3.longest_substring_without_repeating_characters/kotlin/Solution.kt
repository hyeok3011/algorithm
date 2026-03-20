// https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun lengthOfLongestSubstring(s: String): Int {
        val letters = BooleanArray(128)
        var maxLength = 0
        var start = 0
        for (i in s.indices) {
            val letterIndex = s[i].code
            if (letters[letterIndex]) {
                while (s[i] != s[start]) {
                    letters[s[start].code] = false
                    start++
                }
                start++
            }
            letters[letterIndex] = true
            maxLength = maxOf(maxLength, i - start + 1)
        }

        return maxLength
    }
}
