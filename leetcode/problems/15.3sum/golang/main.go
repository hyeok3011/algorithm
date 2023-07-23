package main

import "sort"

// https://leetcode.com/problems/3sum/
// Runtime 45 ms Beats 87.40% Memory 7.4 MB Beats 94.26%
// 문제를 제대로 이해했는지 확인겸 3중 for돌렸다가 바로 timeout컷 당해버림
// 이후 two pointer로 접근해서 다시 품
func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		// for prevent duplication
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// for prevent duplication
				left += 1
				right -= 1
				for left < right && nums[left] == nums[left-1] {
					left += 1
				}
				for left < right && nums[right] == nums[right+1] {
					right = -1
				}
			} else if sum > 0 {
				right -= 1
			} else {
				left += 1
			}
		}

	}
	return result
}

func threeSum2(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	checkDuplicated := make(map[int]map[int]bool)
	for i := 0; i < len(nums)-2; i++ {
		if _, ok := checkDuplicated[nums[i]]; ok {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if _, ok := checkDuplicated[nums[i]][nums[left]]; ok {
				left += 1
				continue
			}
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				checkDuplicated[nums[i]] = make(map[int]bool)
				checkDuplicated[nums[i]][nums[left]] = true
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left += 1
				right -= 1
			} else if sum > 0 {
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return result
}
