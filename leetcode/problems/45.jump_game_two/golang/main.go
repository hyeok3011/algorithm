package main

// https://leetcode.com/problems/jump-game-ii/description/
// Runtime 9 ms Beats 91.21% Memory 5.9 MB Beats 93.3%
// 처음에 문제를 제대로 이해하지 못하고 시간복잡도를 O(n^2)로 만들었다가
// 천천히 문제를 보며 다시 풀어 O(n)으로 변경하였으나 Leetcode문제인건지 속도나 큰 차이가 없음
// -0-
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	steps := 0
	currentMaxIndex := 0
	maxReach := 0

	for i := 0; i < len(nums)-1; i++ {
		if maxReach < i+nums[i] {
			maxReach = i + i + nums[i]
		}

		if i == currentMaxIndex {
			steps++
			currentMaxIndex = maxReach
		}
	}

	return steps
}

func jump2(nums []int) int {
	step := 0
	currentIndex := 0
	for currentIndex < len(nums)-1 {
		step += 1
		maxJumpIndex := currentIndex + nums[currentIndex]
		if maxJumpIndex >= len(nums)-1 {
			break
		}
		for i := currentIndex; i <= currentIndex+nums[currentIndex]; i++ {
			if maxJumpIndex+nums[maxJumpIndex] < i+nums[i] {
				maxJumpIndex = i
			}
		}
		currentIndex = maxJumpIndex
	}

	return step

}
