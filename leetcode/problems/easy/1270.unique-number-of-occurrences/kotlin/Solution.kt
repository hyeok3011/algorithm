
// https://leetcode.com/problems/unique-number-of-occurrences/?envType=study-plan-v2&envId=leetcode-75
class Solution {
    fun uniqueOccurrences(arr: IntArray): Boolean {
        val occurrencesMap = mutableMapOf<Int, Int>()
        for (num in arr) {
            occurrencesMap[num] = occurrencesMap.getOrDefault(num, 0) + 1
        }

        val occurrencesSet = mutableSetOf<Int>()
        for (count in occurrencesMap.values) {
            if (occurrencesSet.contains(count)) {
                return false
            }
            occurrencesSet.add(count)
        }
        return true
    }
}