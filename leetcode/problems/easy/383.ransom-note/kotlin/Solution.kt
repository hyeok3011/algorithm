// https://leetcode.com/problems/ransom-note/?envType=problem-list-v2&envId=rab78cw1
class Solution {
    fun canConstruct(ransomNote: String, magazine: String): Boolean {
        val letters = IntArray(26)
        for (char in ransomNote) {
            letters[char - 'a']++
        }

        for (char in magazine) {
            letters[char - 'a']--
        }

        return letters.all { it <= 0 }
    }

    fun canConstruct2(ransomNote: String, magazine: String): Boolean {
        val letters = IntArray(26)
        var magazineIndex = 0
        for (char in ransomNote) {
            val letterIndex = char - 'a'
            while(letters[letterIndex] == 0) {
                if (magazineIndex >= magazine.length) return false
                letters[magazine[magazineIndex] - 'a']++
                magazineIndex++
            }
            letters[letterIndex]--
        }


        return true
    }   
}
