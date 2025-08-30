package main

import (
	"sort"
)

// @@@@@@
// 문제를 제대로 읽지 않고 무식하고 permutation문제인줄 알고 고생만함.
// 이 문제는 subset문제임  문제가 많이 단순회되었음.. 괜히 이상한 방향으로 고생만 했다.
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	currentSubset := []int{}

	findSubsetsRecursive(&result, currentSubset, nums, 0)

	return result
}

func findSubsetsRecursive(result *[][]int, currentSubset []int, nums []int, start int) {
	subsetCopy := make([]int, len(currentSubset))
	copy(subsetCopy, currentSubset)
	*result = append(*result, subsetCopy)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		currentSubset = append(currentSubset, nums[i])

		findSubsetsRecursive(result, currentSubset, nums, i+1)

		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}

// 재귀를 사용하지 않고 깔끔하게 푼  이 방법이 더 좋은듯
func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)
	subsets := [][]int{[]int{}}
	prevAdditionsStartIndex := 0
	for i := 0; i < len(nums); i++ {
		startIndex := 0
		if i > 0 && nums[i] == nums[i-1] {
			startIndex = prevAdditionsStartIndex
		}
		size := len(subsets)
		for j := startIndex; j < size; j++ {
			newSubset := append([]int{}, subsets[j]...)
			newSubset = append(newSubset, nums[i])
			subsets = append(subsets, newSubset)
		}
		prevAdditionsStartIndex = size
	}
	return subsets
}
