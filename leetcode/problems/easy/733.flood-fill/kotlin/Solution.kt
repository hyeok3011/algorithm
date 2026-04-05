// https://leetcode.com/problems/flood-fill/description/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun floodFill(image: Array<IntArray>, sr: Int, sc: Int, color: Int): Array<IntArray> {
        val originColor = image[sr][sc]

        fun fill(row: Int, col: Int) {
            if (row !in image.indices || col !in image[row].indices) return
            if (image[row][col] != originColor || image[row][col] == color) return

            image[row][col] = color
            fill(row + 1, col)
            fill(row - 1, col)
            fill(row, col + 1)
            fill(row, col - 1)
        }

        fill(sr, sc)
        return image
    }
}
