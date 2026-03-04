// https://leetcode.com/problems/power-of-four/
// @@@@@@ 힌트보고 힘들게 풀은 문제
// 3가지 조건이 존재한다.
// 조건1: 0보다 큰 수
// 조건2: 2의 거듭제곱은 비트에서 1이 딱 하나만 존재할 수 있다.(2의 거듭제곱인지 4의 거듭제곱인지 모르나 범위를 좁힌다.)
// 조건3: 4의 거듭제곱은 항상 3으로 나누면 나머지가 1 이다 
class Solution {
    fun isPowerOfFour(n: Int): Boolean {
        return n > 0 && (n and (n - 1)) == 0 && (n - 1) % 3 == 0
    }

    // 다른 사람들은 비트 위치를 이용해서 풀었다.
    fun AnotherSolution(n: Int): Boolean {
        return n > 0 && (n and (n - 1)) == 0 && n and 0x55555555 != 0
    }
}