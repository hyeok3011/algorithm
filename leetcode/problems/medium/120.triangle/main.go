package main

/*
https://leetcode.com/problems/triangle/
피라미드 꼭대기에서 맨 아래층까지 최소 비용으로 내려올 수 있는 경로를 찾으라는 뜻이다.
완전탐색으로 풀어도 충분히 풀리겠지만 DP로 푸는것이 더 효율적일듯 싶다.

문제에서 공간 복잡도를 O(n)으로 수행해서 푸는 방향으로 지시했다. n은 총 행의 수다
처음에는 dp를 n크기 만큼 만들어 계속 비교하는 방식으로 했다.
로직도 조금 더럽고 O(2n)을 사용할수밖에 없는 상황이였다.

그래서 조금 생각을 바꿨다.
dp에 따로 데이터를 관리하는것보다 triangle자체를 사용하여 문제를 푸는것이다.
위 방법으로 풀면 메모리를 더 적게 사용할 수 있다.

Runtime 0 ms Beats 100% Memory 2.9 MB Beats 95.36%
DP문제는 항상 어렵다
*/

func minimumTotal(triangle [][]int) int {
	var minVal int
	for i := len(triangle) - 1; i > 0; i-- {
		for j := 0; j < len(triangle[i])-1; j++ {
			minVal = min(triangle[i][j], triangle[i][j+1])
			triangle[i-1][j] += minVal
		}

	}
	return triangle[0][0]
}

func min(data ...int) int {
	minVal := data[0]

	for _, v := range data {
		if minVal > v {
			minVal = v
		}
	}

	return minVal
}
