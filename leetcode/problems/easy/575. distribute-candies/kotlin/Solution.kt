// https://leetcode.com/problems/distribute-candies/description/
class Solution {
    fun distributeCandies(candyType: IntArray): Int {
        val set = mutableSetOf<Int>()
        val maxCandyCount = candyType.size / 2
        for (candy in candyType) {
            set.add(candy)
            if (set.size == maxCandyCount) {
                return maxCandyCount
            }
        }

        return set.size
    }
}
