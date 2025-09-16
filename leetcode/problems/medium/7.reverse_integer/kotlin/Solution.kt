// https://leetcode.com/problems/reverse-integer/description/

class Solution {
    fun reverse(x: Int): Int {
        var answer:Long = 0
        var number = x.toLong() * x.sign
        while (number > 0) {
            answer = answer * 10 + number % 10
            if (answer > Int.MAX_VALUE) {
                return 0
            }
            number /= 10
        }

        return answer.toInt() * x.sign
    }

    fun reverseString(x: Int): Int {
        return try {
            (x * x.sign).toString().reversed().toInt() * x.sign
        } catch (_: NumberFormatException) {
            0
        }
    }
}