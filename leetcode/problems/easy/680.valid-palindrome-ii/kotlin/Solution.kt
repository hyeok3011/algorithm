// @@@ kotlin의 경우 CharArray를 사용하면 조금 더 효율적인, Golang과 다른게 String타입에 대한 관리 방법이 다름.
class Solution {
    fun isPalindrome(s: CharArray, l: Int, r: Int): Boolean{
        var left = l
        var right = r
        while (left < right) {
            if (s[left] != s[right]) return false
            left++
            right--
        }
        return true
    }
    fun validPalindrome(s: String): Boolean {
        var l = 0
        var r = s.length-1
        val arr = s.toCharArray()
        while (l < r) {
            if (arr[l] != arr[r]) {
                return isPalindrome(arr, l+1,r) || isPalindrome(arr, l, r-1)
            }
            l++
            r--
        }
        return true
    }
}
