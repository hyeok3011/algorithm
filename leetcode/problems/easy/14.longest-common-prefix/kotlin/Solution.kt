// https://leetcode.com/problems/longest-common-prefix/description/
class Solution {
    fun longestCommonPrefix(strs: Array<String>): String {
        var prefix = strs[0]
        var prefixLen = prefix.length

        for (i: Int in 1..strs.size - 1) {
            for (j: Int in 0..prefixLen - 1) {
                if (strs[i].length <= j  ||strs[i][j] != prefix[j]) {
                    prefixLen = j
                    break
                }
            }
        }
        if (prefixLen == -1) {
            return ""
        }
        return prefix.substring(0, prefixLen)
    }

    fun longestCommonPrefix2(strs: Array<String>): String {
        return strs.reduceOrNull { acc, s -> 
            acc.commonPrefixWith(s) 
        } ?: ""
    }
}
