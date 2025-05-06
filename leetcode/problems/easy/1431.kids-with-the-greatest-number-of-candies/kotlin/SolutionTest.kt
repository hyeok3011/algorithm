package problems.kids_with_candies

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

internal class SolutionTest {

    private val solution = Solution()

    @Test
    fun `case`() {
        val candies = intArrayOf(2, 3, 5, 1, 3)
        val extraCandies = 3
        val expected = listOf(true, true, true, false, true)
        assertEquals(expected, solution.kidsWithCandies(candies, extraCandies)) 
    }
}