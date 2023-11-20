package main

// https://school.programmers.co.kr/learn/courses/30/lessons/43165
// 간단한 DFS BFS문제 이다.
// 문제의 허들이 낮아 메모이 제이션을 진행하지 않아도 충분히 통과 되었다.
// index기준으로 같은 sum값을 가지고 있으면 이후 연산도 같은 결과가 나올거라 판단하여 풀어봤다.
// 대충 벤치마크를 돌려보니 depth가 길어지면 메모이제이션을 하지 않은 코드는 속도가 느려지는것을 확인했다.
// 끗
func solution(numbers []int, target int) int {
	var dfs func(int, int)

	var totalCount int
	dfs = func(index, sum int) {
		if index == len(numbers) {
			if sum == target {
				totalCount++
			}
			return
		}

		dfs(index+1, sum+numbers[index])
		dfs(index+1, sum-numbers[index])
	}

	dfs(0, 0)
	return totalCount
}

func solution2(numbers []int, target int) int {
	memoization := make(map[[2]int]int)
	var dfs func(int, int) int

	dfs = func(index, sum int) int {
		if index == len(numbers) {
			if sum == target {
				return 1
			}
			return 0
		}

		if val, found := memoization[[2]int{index, sum}]; found {
			return val
		}

		memoization[[2]int{index, sum}] = dfs(index+1, sum+numbers[index]) + dfs(index+1, sum-numbers[index])
		return memoization[[2]int{index, sum}]
	}

	return dfs(0, 0)
}
