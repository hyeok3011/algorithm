// https://leetcode.com/problems/longest-palindromic-substring/description/?envType=problem-list-v2&envId=rab78cw1
// 흠... Golang버전이 더 나을듯함. 이유는 Golang에서는 다음 인덱스에 같은 숫자가 있다면 미리 밀어넣고 시작하는 개념이 있어서 다음 인덱스 탐색을 시작하기전에 동일한지
// 검사하고 continue하는 방향으로 최적화가 가능할것같음.
class Solution {
    fun longestPalindrome(s: String): String {
        if (s.length <= 1) return s

        var start = 0
        var end = 0

        for (i in s.indices) {
            val (oddStart, oddEnd) = findPalindromeRange(s, i, i)
            if (oddEnd - oddStart > end - start) {
                start = oddStart
                end = oddEnd
            }

            val (evenStart, evenEnd) = findPalindromeRange(s, i, i + 1)
            if (evenEnd - evenStart > end - start) {
                start = evenStart
                end = evenEnd
            }
        }

        return s.substring(start, end + 1)
    }

    private fun findPalindromeRange(s: String, left: Int, right: Int): Pair<Int, Int> {
        var l = left
        var r = right
        while (l >= 0 && r < s.length && s[l] == s[r]) {
            l--
            r++
        }
        return (l + 1) to (r - 1)
    }
}
