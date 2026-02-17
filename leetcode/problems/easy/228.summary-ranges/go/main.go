package main

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	answer := []string{}
	i := 0
	var buf strings.Builder
	for i < len(nums) {
		start := i
		for i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
			i++
		}

		buf.WriteString(strconv.Itoa(nums[start]))
		if start != i {
			buf.WriteString("->")
			buf.WriteString(strconv.Itoa(nums[i]))
		}
		answer = append(answer, buf.String())
		buf.Reset()
		i++
	}
	return answer
}
