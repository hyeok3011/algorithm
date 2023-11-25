package main

import "sort"

/*
*
https://school.programmers.co.kr/learn/courses/30/lessons/42861
최소 비용으로 섬끼리 연결해야 한다는것은 비용이 저렴한 선을 사용하며 선이 적을수록 좋다는 말이기 때문이다.
즉 싸이클이 없다 라고 판단했다.
싸이클이 없는 그래프? 최소 연결? ->minumum spanning tree? 라는 의식의 흐름으로 이어젔다.
개념만 알고있었고 사실 실제로 사용해본적은 없어 찾아봤다.
https://en.wikipedia.org/wiki/Minimum_spanning_tree
그래프의 모든 노드를 포함하고 사이클을 형성하지 않는 부분 그래프 중에서, 간선의 가중치 합이 최소인 트리
사이클을 형성하지 않고 그래프의 모든 노드를 포함하며 간선의 가중치 합이 최소인 트리 라고 소개한느거 보면 제대로 찾아온것 같다.
1. Borůvka's algorithm
2. Prim's algorithm
3. Kruskal's algorithm
4. reverse-delete algorithm
위에서 소개한 4가지 알고리즘 방법이다.
prim이나 Kruskal는 많이 들어본 알고리즘 들이 나왔다.
위 알고리즘을 모두 이해하지는 못했고 먼저 유명한 크루스칼 알고리즘 먼저 써봤다.
https://en.wikipedia.org/wiki/Kruskal%27s_algorithm, https://en.wikipedia.org/wiki/Disjoint-set_data_structure
자료들 보면서 어떻게 풀기는 했는데 위 내용들을 모두 정리하는 시간을 가져야 겠다.
*
*/
func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent []int, a int, b int) {
	a = find(parent, a)
	b = find(parent, b)
	if a < b {
		parent[b] = a
	} else {
		parent[a] = b
	}
}

func solution(n int, costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})

	// disjointset
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	totalCost := 0
	for _, cost := range costs {
		a, b, c := cost[0], cost[1], cost[2]

		// 사이클을 생성하지 않기위해 확인
		if find(parent, a) != find(parent, b) {
			union(parent, a, b)
			totalCost += c
		}
	}
	return totalCost
}
