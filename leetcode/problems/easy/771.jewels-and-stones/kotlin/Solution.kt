// https://leetcode.com/problems/jewels-and-stones/submissions/1943904532/

class Solution {
    fun numJewelsInStones(jewels: String, stones: String): Int {
        val jewelSet = BooleanArray(128)
        for (c in jewels) {
            jewelSet[c.code] = true
        }

        var count = 0
        for (c in stones) {
            if (jewelSet[c.code]) count++
        }
        return count
    }
}