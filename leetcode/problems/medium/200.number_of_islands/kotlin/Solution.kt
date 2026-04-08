// https://leetcode.com/problems/number-of-islands/description/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun numIslands(grid: Array<CharArray>): Int {
        var islandCount = 0
        fun markIsland(row: Int, col: Int) {
            if (row !in grid.indices || col !in grid[row].indices) return
            if (grid[row][col] != '1') return

            grid[row][col] = '0'
            markIsland(row - 1, col)
            markIsland(row + 1, col)
            markIsland(row, col - 1)
            markIsland(row, col + 1)
        }

        for (row in 0 until grid.size) {
            for (col in 0 until grid[row].size) {
                if (grid[row][col] == '1') {
                    markIsland(row, col)
                    islandCount++
                }
            }
        }

        return islandCount
    }
}
