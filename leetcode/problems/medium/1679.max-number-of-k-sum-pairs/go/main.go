package main

import "slices"

// https://leetcode.com/problems/max-number-of-k-sum-pairs/description/?envType=study-plan-v2&envId=leetcode-75
func maxOperations(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}

	cnt := 0
	numberCount := make(map[int]int)
	for i := range nums {
		if numberCount[k-nums[i]] > 0 {
			numberCount[k-nums[i]]--
			cnt++
		} else {
			numberCount[nums[i]]++
		}
	}
	return cnt
}

func maxOperationsTwoPointer(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}

	slices.Sort(nums)
	left, right := 0, len(nums)-1
	cnt := 0
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			cnt++
			right--
			left++
		} else if sum > k {
			right--
		} else {
			left++
		}
	}
	return cnt
}
