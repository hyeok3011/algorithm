package main

// https://school.programmers.co.kr/learn/courses/30/lessons/161990
func solution(wallpaper []string) []int {
	minLeft, maxRight, minTop, maxBottom := len(wallpaper[0]), 0, len(wallpaper), 0
	for i := range wallpaper {
		for j := range wallpaper[i] {
			if wallpaper[i][j] == '#' {
				if minLeft > j {
					minLeft = j
				}
				if maxRight <= j {
					maxRight = j + 1
				}

				if minTop > i {
					minTop = i
				}

				if maxBottom <= i {
					maxBottom = i + 1
				}
			}
		}
	}
	return []int{minTop, minLeft, maxBottom, maxRight}
}
