// https://leetcode.com/problems/word-break/description/?envType=problem-list-v2&envId=rab78cw1
// @@@@ 이러한 문제에서 항상 느끼지만 처음부터 DP로 접근하는것이 너무 어려움. 과거 선택이 미래 선택에 영향을 안주기때문에
// 여기까지 도달이 가능한가???? 그리고 이후 판단이 충분히 가능한가??? 이러한 의식의 흐름을 잘 흘러가지 않음...ㅜ

class Solution {
    fun wordBreakRecursion(s: String, wordDict: List<String>): Boolean {
        val wordSet = wordDict.toHashSet()
        val visited = BooleanArray(s.length)
        val maxWordLen = wordDict.maxOf{ it.length}
        fun recursion(start: Int): Boolean {
            if (start >= s.length) return true
            if (visited[start]) return false
            val end = minOf(s.length, start + maxWordLen)
            for (i in start until end) {
                if (s.substring(start, i + 1) in wordSet && recursion(i + 1)) {
                    return true
                }
            }
            visited[start] = true
            return false
        }

        return recursion(0)
    }
    
    fun wordBreak(s: String, wordDict: List<String>): Boolean {
        val wordSet = wordDict.toHashSet()
        val dp = BooleanArray(s.length + 1)
        val maxWordLen = wordDict.maxOf{it.length}
        dp[0] = true
        for (i in 1 .. s.length) {
            for (j in i downTo max(0, i - maxWordLen)) {
                if (dp[j] && s.substring(j, i) in wordSet) {
                    dp[i] = true
                    break
                }
            }
        }
        return dp[s.length]
    }
}
