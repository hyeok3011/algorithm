package main

// https://leetcode.com/problems/combinations/description/
// Runtime 121 ms Beats 16.75% Memory 50.4 MB Beats 16.45%
// 넘나 느린것.... 0ms는 어떻게 나온건지 모르겠음...
// @@@@@@@@@@@@@@@@@2 다시 풀어볼것
func combine(n int, k int) [][]int {
	combinations := [][]int{}
	current := make([]int, 0, k)
	recursion(current, 1, n, k, &combinations)
	return combinations
}

func recursion(current []int, start, end, k int, combinations *[][]int) {
	if k == len(current) {
		combination := make([]int, len(current))
		copy(combination, current)
		*combinations = append(*combinations, combination)
		return
	}

	for i := start; i <= end-(k-len(current))+1; i++ {
		current = append(current, i)
		recursion(current, i+1, end, k, combinations)
		current = current[:len(current)-1]
	}
}
