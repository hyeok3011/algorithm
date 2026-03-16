// https://leetcode.com/problems/decode-string/description/
class Solution {
    fun decodeString(s: String): String {
        var index = 0

        fun decode(): String {
            var result = ""
            var repeat = 0

            while (index < s.length) {
                when {
                    s[index].isDigit() -> {
                        repeat = repeat * 10 + (s[index] - '0')
                        index++
                    }
                    s[index] == '[' -> {
                        index++
                        val decoded = decode()
                        result += decoded.repeat(repeat)
                        repeat = 0
                    }
                    s[index] == ']' -> {
                        index++
                        return result
                    }
                    else -> {
                        result += s[index]
                        index++
                    }
                }
            }

            return result
        }

        return decode()
    }
}