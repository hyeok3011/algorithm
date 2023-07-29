package main

import "sort"

// https://leetcode.com/problems/permutations-ii/
// Runtime 4 ms Beats 68.29% Memory 3.7 MB Beats 76.83%
func permuteUnique(nums []int) [][]int {
	permutations := [][]int{}
	used := make([]bool, len(nums))
	current := []int{}

	sort.Ints(nums)
	generatePermutations(nums, current, &permutations, used)
	return permutations
}

func generatePermutations(nums, current []int, permutations *[][]int, used []bool) {
	if len(current) == len(nums) {
		permutation := make([]int, len(nums))
		copy(permutation, current)
		(*permutations) = append((*permutations), permutation)
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && !used[i-1] && nums[i] == nums[i-1] {
			continue
		}

		current = append(current, nums[i])
		used[i] = true
		generatePermutations(nums, current, permutations, used)
		used[i] = false
		current = current[:len(current)-1]
	}
}

func permuteUnique2(nums []int) [][]int {
	res := [][]int{}
	findPermutations(nums, 0, &res)
	return res
}

func findPermutations(nums []int, idx int, res *[][]int) {
	if idx == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		(*res) = append(*res, temp)
		return
	}

	used := make(map[int]bool)
	for i := idx; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		nums[i], nums[idx] = nums[idx], nums[i]
		findPermutations(nums, idx+1, res)
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
