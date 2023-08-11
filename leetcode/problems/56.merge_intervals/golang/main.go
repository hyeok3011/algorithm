package main

import "sort"

// https://leetcode.com/problems/merge-intervals/
// Runtime 14 ms Beats 97.57% Memory 6.8 MB Beats 72.80%
// @@@@@@@@@@@
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	index := 1

	for index <= len(intervals)-1 {
		if intervals[index-1][1] >= intervals[index][0] {
			intervals[index-1][1] = max(intervals[index-1][1], intervals[index][1])
			intervals = append(intervals[:index], intervals[index+1:]...)
		} else {
			index += 1
		}
	}

	return intervals
}

// 다른사람의 풀이를 보고 다시 짜봤다. slice를 잘라내지 않는 풀이 방법이다.
func merge2(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[mergedIndex][1] >= intervals[i][0] {
			intervals[mergedIndex][1] = max(intervals[mergedIndex][1], intervals[i][1])
		} else {
			mergedIndex += 1
			intervals[mergedIndex] = intervals[i]
		}
	}

	return intervals[:mergedIndex+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
