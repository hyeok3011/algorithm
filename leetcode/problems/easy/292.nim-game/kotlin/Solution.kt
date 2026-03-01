// https://leetcode.com/problems/nim-game/description/
// 게임이론인지 모르고 그냥 풀어버림
// 내 차례에 4의 배수 → 나는 뭘 해도 상대에게 4의 배수가 아닌 숫자를 줌 → 상대가 다시 나에게 4의 배수를 줌 → ... → 결국 내가 4를 받게 됨 → loss
class Solution {
    fun canWinNim(n: Int): Boolean {
        if (n <= 3) return true
        return n % 4 != 0
    }
}