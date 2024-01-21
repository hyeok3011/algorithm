package main

/*
https://leetcode.com/problems/subsets/description/
간단한 조합 만드는 문제이다.
generateCombinations에서 N을 따로 받지 않고 한번에 만드는 방법도 있겠지만 헷갈려서
그냥 아래 방법으로 풀었다.
generateCombinations for순회에서 len(nums)-(n-depth)+1 조건을 사용하여
의미없는 콜스택을 줄였다.
*/

func subsets(nums []int) [][]int {
	subsets := [][]int{}
	for i := 0; i <= len(nums); i++ {
		generateCombinations(nums, i, 0, 0, make([]int, i), &subsets)
	}
	return subsets
}

func generateCombinations(nums []int, n, startIndex, depth int, current []int, combinations *[][]int) {
	if depth == n {
		combination := make([]int, n)
		copy(combination, current)
		*combinations = append(*combinations, combination)
		return
	}

	for i := startIndex; i < len(nums)-(n-depth)+1; i++ {
		current[depth] = nums[i]
		generateCombinations(nums, n, i+1, depth+1, current, combinations)
	}
}
