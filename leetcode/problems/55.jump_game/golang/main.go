package main

// https://leetcode.com/problems/jump-game/
// Runtime 42 ms Beats 100% Memory 7.2 MB Beats 85.2%
// 음... bottom up? 방식으로 올라가서 처리 시간 복잡도O(n)
func canJump(nums []int) bool {
	targetIndex := len(nums) - 1
	currentIndex := targetIndex - 1
	for currentIndex >= 0 {
		if nums[currentIndex] == 0 {
			currentIndex -= 1
			continue
		}

		if (nums[currentIndex] + currentIndex) >= targetIndex {
			targetIndex = currentIndex
		}
		currentIndex -= 1
	}

	if targetIndex != 0 {
		return false
	}

	return true
}

func AnotherSolution(nums []int) bool {
	maxReachableIndex := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReachableIndex {
			return false
		}
		maxReachableIndex = max(maxReachableIndex, i+nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 실패한 알고리즘
func backtracking(nums []int, index int) bool {
	if index == len(nums)-1 {
		return true
	}

	if nums[index] == 0 {
		return false
	}

	for i := nums[index]; i >= 1; i-- {
		if (index + i) > len(nums)-1 {
			continue
		}
		if fail(nums, index+i) {
			return true
		}
	}

	return false
}
