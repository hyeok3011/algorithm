package main

// https://leetcode.com/problems/container-with-most-water/description/
// Runtime 81 ms Beats 80.98% Memory 8.2 MB Beats 52.32%
// Two Pointer
// 제대로 푼거같은데 너무 느림... 뭐지?
func maxArea(height []int) int {
	maximum := 0
	start, end := 0, len(height)-1

	for start != end {
		containerWidth := end - start
		containerHeight := min(height[start], height[end])
		area := containerWidth * containerHeight

		if height[start] > height[end] {
			end -= 1
		} else {
			start += 1
		}
		maximum = max(maximum, area)
	}
	return maximum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
