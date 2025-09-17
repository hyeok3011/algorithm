// https://leetcode.com/problems/integer-to-roman/description/
class Solution {
    fun intToRoman(num: Int): String {
        val romanValues = listOf(
            1000 to "M", 900 to "CM", 500 to "D", 400 to "CD", 100 to "C",
            90 to "XC", 50 to "L", 40 to "XL", 10 to "X",
            9 to "IX", 5 to "V", 4 to "IV", 1 to "I"
        )

        val result = StringBuilder()
        var remaining = num

        romanValues.forEach { (value, symbol) ->
            while (remaining >= value) {
                result.append(symbol)
                remaining -= value
            }
        }

        return result.toString()
    }
}

// 어이가 없는 방법 ㅋㅋㅋ
val symbols = arrayOf(
    1000 to "M",
    900 to "CM",
    500 to "D",
    400 to "CD",
    100 to "C",
    90 to "XC",
    50 to "L",
    40 to "XL",
    10 to "X",
    9 to "IX",
    5 to "V",
    4 to "IV",
    1 to "I",
)

@JvmField
val numerals = Array(4000) { intToRoman(it) }

fun intToRoman(num: Int): String = buildString {
    symbols.fold(num) { acc, (n, str) ->
        val times = acc / n
        for (i in 0..<times) append(str)
        acc - n * times
    }
}

fun Any?.intToRoman(num: Int): String {
    return numerals[num]
}
inline fun Solution(): Int = 0