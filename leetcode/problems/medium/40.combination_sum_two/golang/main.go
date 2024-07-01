package main

import "sort"

// https://leetcode.com/problems/combination-sum-ii/description/
// Runtime 0 ms Beats 100% Memory 2.6 MB Beats 84.65%
func combinationSum2(candidates []int, target int) [][]int {
	combination := []int{}
	answer := [][]int{}
	sort.Ints(candidates)
	used := make([]bool, len(candidates))
	generateCombination2(0, target, candidates, &combination, &answer, used)
	return answer
}

func generateCombination(index, target int, candidates []int, combination *[]int, result *[][]int, used []bool) {
	if (*combination)[0] > target {
		return
	}

	if (*combination)[0] == target {
		(*result) = append((*result), append([]int{}, (*combination)[1:]...))
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > 0 && !used[i-1] && candidates[i] == candidates[i-1] {
			continue
		}
		(*combination)[0] += candidates[i]
		(*combination) = append((*combination), candidates[i])
		used[i] = true
		generateCombination(i+1, target, candidates, combination, result, used)
		(*combination)[0] -= candidates[i]
		(*combination) = (*combination)[:len((*combination))-1]
		used[i] = false
	}
}

func generateCombination2(index, target int, candidates []int, combination *[]int, result *[][]int, used []bool) {
	if target < 0 {
		return
	}

	if target == 0 {
		(*result) = append((*result), append([]int{}, (*combination)...))
		return
	}

	for i := index; i < len(candidates); i++ {
		if i > 0 && !used[i-1] && candidates[i] == candidates[i-1] {
			continue
		}

		(*combination) = append((*combination), candidates[i])
		used[i] = true

		generateCombination2(i+1, target-candidates[i], candidates, combination, result, used)

		(*combination) = (*combination)[:len((*combination))-1]
		used[i] = false
	}
}
