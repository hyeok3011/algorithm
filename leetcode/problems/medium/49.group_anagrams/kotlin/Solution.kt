// https://leetcode.com/problems/group-anagrams/
class Solution {
    fun groupAnagrams(strs: Array<String>): List<List<String>> {
        val strGroup = mutableMapOf<String, MutableList<String>>()
        for (i in strs.indices) {
            val letters = IntArray(26)
            val currentStr = strs[i]
            for (j in 0 until currentStr.length) {
                letters[currentStr[j] - 'a']++
            }
            val key = letters.joinToString(",")
            strGroup.getOrPut(key) { mutableListOf() }.add(currentStr)
        }
        val result = mutableListOf<MutableList<String>>()
        return strGroup.values.toList()
    }
}
