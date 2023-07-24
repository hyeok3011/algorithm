package main

import (
	"sort"
)

// https://leetcode.com/problems/3sum-closest/description/
// Runtime 9 ms Beats 89.4 Memory 3.2 MB Beats 95.77%
func threeSumClosest(nums []int, target int) int {
	closest := 9999999999
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// for prevent duplication
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				right -= 1
			} else {
				left += 1
			}
			if getDistance(target, sum) < getDistance(target, closest) {
				closest = sum
			}
		}

	}
	return closest
}

func getDistance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
