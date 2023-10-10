package main

const (
	SHEEP = 0
	WOLF  = 1
)

// https://school.programmers.co.kr/learn/courses/30/lessons/92343
//

// 최대한의 양을 모으는 문제다.
// 양의 개수가 늑대의 개수보다 많을경우 노드를 이동이 가능하다.
// 문제는 보고 완전 탐색 문제라고 생각했다.
// 하지만 오른쪽으로 이동하다가 다시 왼쪽으로 이동해야 하는 상황이 존재한다.
// 단순하게 짜먄되겠지라는 생각을 했다가 바로 실패했다.
// 그래서 계속 고민하던중 https://school.programmers.co.kr/questions/25736에서 힌트를 얻었다.
// 이미 방문한 노드는 제외해놓고 방문 가능한 노드를 계속 순회하는 방법이다.
// 이 방법 외에 solution.py에 bitmasking방법도 있는데 이것도 기가막힌다....
func solution(info []int, edges [][]int) int {
	graph := make([][]int, len(info))
	for i := range graph {
		graph[i] = []int{}
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	maxSheep := 0

	var dfs func(cur int, sheepCount int, wolfCount int, available []bool)
	dfs = func(cur int, sheepCount int, wolfCount int, available []bool) {
		if info[cur] == SHEEP {
			sheepCount++
		} else {
			wolfCount++
		}

		if sheepCount > maxSheep {
			maxSheep = sheepCount
		}

		if wolfCount >= sheepCount {
			return
		}

		newAvailable := make([]bool, len(available))
		copy(newAvailable, available)
		for _, next := range graph[cur] {
			newAvailable[next] = true
		}

		for i, canVisit := range newAvailable {
			if canVisit {
				newAvailable[i] = false
				dfs(i, sheepCount, wolfCount, newAvailable)
				newAvailable[i] = true
			}
		}
	}

	startAvailable := make([]bool, len(info))
	dfs(0, 0, 0, startAvailable)

	return maxSheep
}
