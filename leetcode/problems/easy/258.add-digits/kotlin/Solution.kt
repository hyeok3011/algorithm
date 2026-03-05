// https://leetcode.com/problems/add-digits/description/
// 이것도 힌트를 보고 힘들게 풀었음. Digital Root이론이라는 개념인듯함, 나중에 정리하면 좋을듯함.
class Solution {
    fun addDigits(num: Int): Int {
        var n = num
        while (n >= 10) {
            n = n.toString().sumOf { it.digitToInt() }
        }
        return n
    }

    fun addDigits2(num: Int): Int {
        if (num == 0) return 0
        return 1 + (num - 1) % 9
    }
}
