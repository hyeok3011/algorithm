package main

import "sort"

// https://school.programmers.co.kr/learn/courses/30/lessons/147354
// 간단하게 정렬후에 xor연산만 하면 될거같아 단순하게 풀어봤다.
// 다른 사람들의 풀이도 크게 다르지 않은듯 하다.

func solution(data [][]int, col int, row_begin int, row_end int) int {
	sort.Slice(data, func(i, j int) bool {
		if data[i][col-1] == data[j][col-1] {
			return data[i][0] > data[j][0]
		}
		return data[i][col-1] < data[j][col-1]
	})

	var hashVal, remainder int

	for i := row_begin; i <= row_end; i++ {
		remainder = 0
		for _, v := range data[i-1] {
			remainder += v % i
		}
		hashVal ^= remainder
	}

	return hashVal
}
