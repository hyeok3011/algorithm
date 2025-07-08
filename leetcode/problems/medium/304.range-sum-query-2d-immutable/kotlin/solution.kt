// https://leetcode.com/problems/range-sum-query-2d-immutable

class NumMatrix(matrix: Array<IntArray>) {

    private val prefixSum: Array<IntArray>

    init {
        prefixSum = calculatePrefixSum(matrix)
    }

    fun sumRegion(row1: Int, col1: Int, row2: Int, col2: Int): Int {
        return prefixSum[row2 + 1][col2 + 1] - prefixSum[row1][col2 + 1] - prefixSum[row2 + 1][col1] + prefixSum[row1][col1]
    }

    private fun calculatePrefixSum(matrix: Array<IntArray>): Array<IntArray> {
        val m = matrix.size
        val n = matrix.getOrNull(0)?.size ?: 0
        
        val sumArray = Array(m + 1) { IntArray(n + 1) }

        for (r in 0 until m) {
            for (c in 0 until n) {
                sumArray[r + 1][c + 1] = sumArray[r][c + 1] + sumArray[r + 1][c] - sumArray[r][c] + matrix[r][c]
            }
        }
        return sumArray
    }
}