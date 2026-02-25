// https://leetcode.com/problems/number-of-1-bits/description/
// 이문제도 최적화 관련내용은 특정 사이즈로 쪼개 캐시처리하는 문제인듯함.
class Solution {
    fun hammingWeight(n: Int): Int {
        var count = 0
        var number = n
        while(number > 0) {
            if (number and 1 == 1) {
                count++
            }
            number = number shr 1
        }
        return count
    }
}
