// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
class Solution {
    fun strStr(haystack: String, needle: String): Int {
        for(i in 0 .. haystack.length - needle.length) {
            if (needle.indices.all { j -> haystack[i + j] == needle[j] }) {
                return i
            }
        }
        return -1
    }
}
