package main

import "sort"

// https://leetcode.com/problems/insert-interval/description/
// 이미 정렬이 되어있는 상태이기 때문에 시작점과 끝점을
func insert(intervals [][]int, newInterval []int) [][]int {
	startIdx := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][1] >= newInterval[0]
	})
	endIdx := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] > newInterval[1]
	})

	if startIdx == endIdx {
		if startIdx == len(intervals) || newInterval[1] < intervals[startIdx][0] {
			result := make([][]int, len(intervals)+1)
			copy(result, intervals[:startIdx])
			result[startIdx] = newInterval
			copy(result[startIdx+1:], intervals[startIdx:])
			return result
		}
	}

	intervals[startIdx][0] = min(intervals[startIdx][0], newInterval[0])
	intervals[startIdx][1] = max(intervals[endIdx-1][1], newInterval[1])
	if diff := endIdx - startIdx; diff > 1 {
		copy(intervals[startIdx+1:], intervals[endIdx:])
		intervals = intervals[:len(intervals)-(diff-1)]
	}
	return intervals
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
