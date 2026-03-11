// https://leetcode.com/problems/jewels-and-stones/submissions/1943904532/

class Solution {
    fun numJewelsInStones(jewels: String, stones: String): Int {
        val jewelSet = BooleanArray(58)
        for (c in jewels) {
            jewelSet[c.code - 65] = true
        }

        var count = 0
        for (c in stones) {
            if (jewelSet[c.code - 65]) count++
        }
        return count
    }
}