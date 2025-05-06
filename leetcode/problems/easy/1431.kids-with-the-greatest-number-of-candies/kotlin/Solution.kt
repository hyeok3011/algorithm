package problems.kids_with_candies

class Solution {
    fun kidsWithCandies(candies: IntArray, extraCandies: Int): List<Boolean> {
        val maxCandyCount: Int = candies.maxOrNull() ?: 0
        
        return candies.map { candyCount ->
            candyCount + extraCandies >= maxCandyCount
        }
    }
}
