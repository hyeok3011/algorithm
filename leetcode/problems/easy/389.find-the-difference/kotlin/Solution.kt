// https://leetcode.com/problems/find-the-difference/
class Solution {
    fun findTheDifference(s: String, t: String): Char {
        val sElem = ByteArray(26)
        val tElem = ByteArray(26)
        for (i in 0 until s.length) {
            sElem[s[i] - 'a']++
            tElem[t[i] - 'a']++
        }
        tElem[t.last() - 'a']++ 

        for (i in 0 until sElem.size) {
            if (sElem[i] != tElem[i]) {
                return ('a' + i)
            }
        }
        return '?'
    }
    // 다른사람의 풀이... XOR연산자 사용...
    fun findTheDifferenceXor(s: String, t: String): Char {
        var sum = 0

        for(i in 0 until s.length){
            sum = sum xor s[i].code
        }

        for(i in 0 until t.length){
            sum = sum xor t[i].code
        }

        return sum.toChar()
    }
}