package problems.easy.countBits

class Solution {
    fun countBits(n: Int): IntArray {
        val answer = IntArray(n + 1)
        for (i: Int in 0..n) {
            var x = i
            var count = 0
            while (x > 0) {
                if (x % 2 == 1) {
                    count += 1
                }
                x = x / 2
            }
            answer.set(i, count)    
        }
        return answer
    }

    // 생각 못할 문제는 아니였으나 도저히 생각이 나지 않음.... 역시 꾸준히 풀어야해... 
    fun countBits2(n: Int): IntArray {
        val answer = IntArray(n + 1)
        for (i in 1..n) {
            answer[i] = answer[i shr 1] + (i and 1)
        }
        return answer
    }
}
