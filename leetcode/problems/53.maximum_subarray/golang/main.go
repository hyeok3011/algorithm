package main

// https://leetcode.com/problems/maximum-subarray/
// Runtime 79 ms Beats 99.52% Memory 8 MB Beats 88.64%
// 개인적으로 어려웠는데 다른사람들의 풀이를 보니 엄청 쉽게풀었다...
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currentSum, maxSum := nums[0], nums[0]
	// maxStartIndex, maxEndIndex := 0, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > currentSum+nums[i] {
			currentSum = nums[i]
			// maxStartIndex = i
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			// maxEndIndex = i
		}
	}
	// log.Println(maxStartIndex, maxEndIndex)
	return maxSum
}

func AnotherSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
