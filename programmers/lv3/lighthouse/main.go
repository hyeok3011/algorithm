package main

// https://school.programmers.co.kr/learn/courses/30/lessons/133500#qna
// 그래프로 데이터를 만들고 우선순위 큐를 사용해서 엣지가 많은경우를 키려고 했지만 실패했다.
// 다음으로 생각한것이 완전탐색이지만 분명 간단한 완전탐색 문제라면 lv3이라고 생각하지 않고 다른방법을 모색했다.
// 결국 선택한 방법은 그리디 방법이다.
// 그리디 방법으로 먼저 깊이 들어간뒤 지역적으로 자식의 노드중에 하나라도 꺼저있으면 반드시 등대를 키는 방식이다.
// 위 방법으로 풀었더니 풀렸다.
// dfs에서 배열들은 어차피 슬라이스의 헤더만을 복사하고 내부 배열 포인터는 유지되기 때문에 값복사 방법을 선택했다.
// 처음부터 익명함수를 쓰면 되지만 알고리즘 문제에서 익명함수 쓰기가 이상하게 싫다.

func solution(n int, lighthouse [][]int) int {
	graph := make([][]int, n+1)
	for _, v := range lighthouse {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	visited := make([]bool, n+1)
	activatedLightHouse := make([]bool, n+1)
	dfs(graph, 1, 1, visited, activatedLightHouse)
	activatedLightHouseCount := 0
	for _, v := range activatedLightHouse {
		if v {
			activatedLightHouseCount++
		}
	}

	return activatedLightHouseCount
}

func dfs(graph [][]int, current, parent int, visited, activatedLightHouse []bool) {
	visited[current] = true
	isLeafNode := true
	isChildLight := true
	for _, next := range graph[current] {
		if !visited[next] {
			dfs(graph, next, current, visited, activatedLightHouse)
			isLeafNode = false

			if !activatedLightHouse[next] {
				isChildLight = false
			}
		}
	}
	if !isLeafNode && !isChildLight {
		activatedLightHouse[current] = true
	}
}
