package main

// https://leetcode.com/problems/permutations/description/
// Runtime 0 ms Beats 100% Memory 2.7 MB Beats 35.80%
// runtine 1ms단위랑 memory 0.1mb가 왔다갔다함...
func permute(nums []int) [][]int {
	permutations := [][]int{}
	used := make([]bool, len(nums))
	current := []int{}
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

		current = append(current, nums[i])
		used[i] = true
		generatePermutations(nums, current, permutations, used)
		used[i] = false
		current = current[:len(current)-1]
	}
}

func permute2(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	cases := [][]int{}
	for i := 0; i < len(nums); i++ {
		temp := append([]int{}, nums[:i]...)
		for _, v := range permute2(append(temp, nums[i+1:]...)) {
			cases = append(cases, append(v, nums[i]))
		}
	}
	return cases
}

// 다른사람의 솔루션 swap을 하며 순열을 구하는 형태
func permute3(nums []int) [][]int {
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

	for i := idx; i < len(nums); i++ {
		if idx == i {
			findPermutations(nums, idx+1, res)
		} else {
			nums[i], nums[idx] = nums[idx], nums[i]
			findPermutations(nums, idx+1, res)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
}
