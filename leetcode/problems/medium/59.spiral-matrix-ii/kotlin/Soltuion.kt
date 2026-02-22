// https://leetcode.com/problems/spiral-matrix-ii/
class Solution {
    fun generateMatrix(n: Int): Array<IntArray> {
        val matrix = Array(n) { _ -> IntArray(n) }
        var left = 0
        var right = n - 1
        var top = 0
        var bottom = n - 1
        var current = 1
        while (left <= right && top <= bottom) {
            for (i in left .. right) {
                matrix[top][i] = current
                current++
            }
            top++
            for (i in top .. bottom) {
                matrix[i][right] = current
                current++
            }
            right--
            for (i in right downTo left) {
                matrix[bottom][i] = current
                current++
            }
            bottom--
            for (i in bottom downTo top) {
                matrix[i][left] = current
                current++
            }
            left++
        }
        return matrix
    }
}
