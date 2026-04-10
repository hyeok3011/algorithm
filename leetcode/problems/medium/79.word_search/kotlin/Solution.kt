// https://leetcode.com/problems/word-search/?envType=problem-list-v2&envId=rab78cw1
// visited를 따로 사용하지 않고 #와같은 다른 Char를 넣어놓고 visited를 확인하면 더 효율성이 높아질듯함.
class Solution {
    fun exist(board: Array<CharArray>, word: String): Boolean {
        val visited = Array(board.size) { BooleanArray(board[0].size) }
        fun dfs(row: Int, col: Int, wordIndex: Int): Boolean {
            if (wordIndex == word.length) return true
            if (row !in board.indices || col !in board[row].indices) return false
            if (visited[row][col]) return false

            if (board[row][col] != word[wordIndex]) {
                return false
            }
            visited[row][col] = true
            if (dfs(row + 1, col, wordIndex + 1) ||dfs(row - 1, col, wordIndex + 1) || dfs(row, col + 1, wordIndex + 1) ||dfs(row, col - 1, wordIndex + 1)) {
                return true
            }
            visited[row][col] = false

            return false
        }

        for (row in 0 until board.size) {
            for (col in 0 until board[row].size) {
                if(board[row][col] == word[0] && dfs(row, col, 0)) {
                    return true
                }
            }
        }
        return false
    }
}