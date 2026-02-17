// https://leetcode.com/problems/different-ways-to-add-parentheses/
// 다른 사람의 깔끔한 코드를 보고 재풀이
class Solution {
    val memo = mutableMapOf<String, List<Int>>()
    fun diffWaysToCompute(expression: String): List<Int> {
        if (expression in memo) {
            return memo.get(expression)!!
        }

        val numbers = mutableListOf<Int>()
        for (i in 0 until expression.length) {
            if (expression[i].isDigit()) {
                continue
            }

            val left = diffWaysToCompute(expression.substring(0, i))
            val right = diffWaysToCompute(expression.substring(i + 1))

            for (l in left) {
                for (r in right) {
                    when (expression[i]) {
                        '+' -> numbers.add(l + r)
                        '-' -> numbers.add(l - r)
                        '*' -> numbers.add(l * r)
                    }
                }
            }

            if (numbers.isEmpty()) {
                numbers.add(expression.toInt())
            }

            memo[expression] = numbers
            return numbers
        }

        return memo[expression] ?: listOf(expression.toInt())
    }
}