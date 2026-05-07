// https://leetcode.com/problems/excel-sheet-column-title/
class Solution {
    fun convertToTitle(columnNumber: Int): String {
            val sb = StringBuilder()
            var number = columnNumber
            while(number > 0) {
                number--
                sb.append('A' + number % 26)
                number = number / 26
            }

            return sb.reverse().toString()
    }
}
