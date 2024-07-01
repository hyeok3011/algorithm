package main

// https://leetcode.com/problems/remove-element/description/
// Runtime 0 ms Beats 100% Memory 2.1 MB Beats 10.29%
func removeElement(nums []int, val int) int {
	validCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[validCount] = nums[i]
			validCount++
		}
	}

	return validCount
}
