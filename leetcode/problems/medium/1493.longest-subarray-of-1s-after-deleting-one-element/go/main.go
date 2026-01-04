package main

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/?envType=study-plan-v2&envId=leetcode-75
// @@@@@ 
func longestSubarray(nums []int) int {
    zeroCount := 0
    left := 0
    maxLen := 0
    for i, v := range nums {
        if v == 0 {
            zeroCount++
        }
        for zeroCount > 1 {
            if nums[left] == 0 {
                zeroCount--
            }
            left++
        }
        maxLen = max(maxLen, i - left)
    }
    return maxLen
}
