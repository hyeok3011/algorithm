package main

// https://leetcode.com/problems/combination-sum/description/
// Runtime 0 ms Beats 100% Memory 2.8 MB Beats 100%

func combinationSum(candidates []int, target int) [][]int {
	combination := []int{0}
	// combination := []int{} // generateCombination2
	answer := [][]int{}
	generateCombination(0, target, candidates, &combination, &answer)
	return answer
}

func generateCombination(index, target int, candidates []int, combination *[]int, result *[][]int) {
	if (*combination)[0] > target {
		return
	}

	if (*combination)[0] == target {
		(*result) = append((*result), append([]int{}, (*combination)[1:]...))
		return
	}

	for i := index; i < len(candidates); i++ {
		(*combination)[0] += candidates[i]
		(*combination) = append((*combination), candidates[i])
		generateCombination(i, target, candidates, combination, result)
		(*combination)[0] -= candidates[i]
		(*combination) = (*combination)[:len((*combination))-1]
	}
}

func generateCombination2(start, target int, candidates []int, combination *[]int, result *[][]int) {
	if target == 0 {
		(*result) = append((*result), append([]int{}, (*combination)...))
		return
	}

	for i := start; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			continue
		}

		(*combination) = append((*combination), candidates[i])
		generateCombination2(i, target-candidates[i], candidates, combination, result)
		(*combination) = (*combination)[:len((*combination))-1]
	}
}
