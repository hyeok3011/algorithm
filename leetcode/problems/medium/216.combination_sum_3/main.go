package main

/*
https://leetcode.com/problems/combination-sum-iii/description/
단순하게 1 ~ 9까지의 숫자중 k개의 숫자를 사용해서 합이 0이 되는 순자를 찾는 조합 문제이다.
sum을 계속 구하지 않고 처음부터 n을 넘겨주며 n-i해서 나중에 0이 되는 조합을 찾았다.
다른 사람들의 풀이도 크게 다르지 않았고
조금 다르게 풀은 사람은 i가 n보다 크면 순회 자체를 하지 않았다.
이런 사소한것 너무 좋다.
*/

func combinationSum3(k int, n int) [][]int {
	validCombinations := [][]int{}
	generateValidCombinations(1, 0, n, make([]int, k), &validCombinations)

	return validCombinations
}

func generateValidCombinations(start, depth, n int, current []int, combinations *[][]int) {
	if depth == len(current) {
		if n == 0 {
			combination := make([]int, len(current))
			copy(combination, current)
			*combinations = append(*combinations, combination)
		}
		return
	}

	for i := start; i <= 9-(len(current)-depth)+1; i++ {
		current[depth] = i
		generateValidCombinations(i+1, depth+1, n-i, current, combinations)
	}
}
