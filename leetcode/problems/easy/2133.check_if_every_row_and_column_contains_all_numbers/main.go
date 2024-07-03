package main

/*
https://leetcode.com/problems/check-if-every-row-and-column-contains-all-numbers/
n * n크기의 보드에서 row와 Col이 각각 1 ~ n까지의 데이터를 가지고있는지 확인하는 문제이다.
풀이 방법은 간단하다.
1 ~ n만큼 있어야하기 때문에 배열을 만들어두고 확인하면 된다.
index는 0부터 시작하기때문에 -1 연산을 할까 패딩을 줄까 하다가 편하게 패딩을 줬다.
끗
*/
func checkValid(matrix [][]int) bool {
	nums := make([]bool, len(matrix)+1)
	for i := range matrix {
		for j := range matrix[i] {
			if nums[matrix[i][j]] {
				return false
			}
			nums[matrix[i][j]] = true
		}

		for j := 0; j < len(matrix); j++ {
			if !nums[matrix[j][i]] {
				return false
			}
			nums[matrix[j][i]] = false
		}
	}
	return true
}
