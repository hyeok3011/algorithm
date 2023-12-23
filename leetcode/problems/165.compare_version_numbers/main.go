package main

import (
	"strconv"
	"strings"
)

/*
https://leetcode.com/problems/compare-version-numbers/
음... 간단하다 dot를 기준으로 나누고 같은 버전 등급별로 비교하면 된다.
등급이란 우리가 사용하는 major, minor, build or patch같은 개념이다.
더 효율적인 방법이 있는데 몰라서 그런건지 간단한 문제처럼 보인다.
*/

func compareVersion(version1 string, version2 string) int {
	reversions1 := strings.Split(version1, ".")
	reversions2 := strings.Split(version2, ".")
	maxLen := max(len(reversions1), len(reversions2))
	for i := 0; i < maxLen; i++ {
		var num1, num2 int
		if i < len(reversions1) {
			num1, _ = strconv.Atoi(reversions1[i])
		}

		if i < len(reversions2) {
			num2, _ = strconv.Atoi(reversions2[i])
		}

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
