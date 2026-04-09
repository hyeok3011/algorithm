// https://leetcode.com/problems/string-to-integer-atoi/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun myAtoi(s: String): Int {
        var index = 0
        while (index < s.length && s[index] == ' ') index++
        if (index >= s.length) return 0

        var postive = true
        if (s[index] == '-' || s[index] == '+') {
            postive = s[index] == '+'
            index++
        }

        val maxDiv10 = Int.MAX_VALUE / 10
        val maxMod10 = Int.MAX_VALUE % 10

        var answer = 0
        while (index < s.length && s[index].isDigit()) {
            val digit = s[index] - '0'

            if (answer > maxDiv10 || (answer == maxDiv10 && digit > maxMod10)) {
                return if (postive) Int.MAX_VALUE else Int.MIN_VALUE
            }

            answer = answer * 10 + digit
            index++
        }

        if (!postive) answer *= -1

        return answer
    }
}