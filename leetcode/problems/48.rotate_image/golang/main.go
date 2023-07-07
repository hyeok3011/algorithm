package main

import "log"

// https://leetcode.com/problems/rotate-image/description/
// Runtime 0 ms Beats 100% Memory 2.4 MB Beats 5.99%

// supported
// Transpose and Reverse by Columns
// 1. Transpose the matrix by swapping elements across the main diagonal.
// 2. Reverse each row of the transposed matrix to achieve the 90-degree rotation.

// 지원받아도 메모리 0.1mb를 최적화 하지 못함....
func rotate(matrix [][]int) {
	if len(matrix) == 1 {
		return
	}

	for _, v := range matrix {
		log.Println(v)
	}

	length := len(matrix)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	log.Println("ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ")
	for _, v := range matrix {
		log.Println(v)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}
