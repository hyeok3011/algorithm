// https://leetcode.com/problems/merge-sorted-array/
class Solution {
    fun merge(nums1: IntArray, m: Int, nums2: IntArray, n: Int): Unit {
        var num1Index = m - 1
        var num2Index = n - 1
        var insertIndex = m + n - 1
        while(num2Index > -1) {
            if (num1Index < 0 || nums1[num1Index] < nums2[num2Index]) {
                nums1[insertIndex] = nums2[num2Index]
                num2Index--
            } else {
                nums1[insertIndex] = nums1[num1Index]
                num1Index--
            }
            insertIndex--
        }

        return
    }
}