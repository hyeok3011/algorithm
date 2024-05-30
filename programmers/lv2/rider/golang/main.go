package main

// https://school.programmers.co.kr/learn/courses/30/lessons/12978
// 1번 마을에서 K시간내로 갈수있는 모든 마을의 횟수를 구하는 문제이다.
// 현재 마을에서 간선이 이어저 있는 마을들을 탐색하며 시간을 구하면 될듯하다.
// 먼저 같은 마을을 가더라도 더 짧은 거리로 갈수있는 방법이 있기때문에 DFS처럼
// 깊이우선 탐색하는 알고리즘은 비효율적일것 같다.
// 먼저 BFS로 풀어봤더니 문제없이 풀렸다.
// 다른 사람들은 다익스트라 알고리즘을 사용해서 푼것같다.
// 다익스트라 알고리즘을 사용하면 인풋값에 따라 더 효율적일수 있겠다라는 생각이 든다.

const (
	MAXIMUM_DISTANCE = 500001
)

func solution(N int, roades [][]int, k int) int {

	villageDistances := make(map[int]int, N)
	roadGraphe := make(map[int]map[int]int, N)
	for i := 1; i <= N; i++ {
		villageDistances[i] = MAXIMUM_DISTANCE
		roadGraphe[i] = make(map[int]int)

	}
	villageDistances[1] = 0

	for _, road := range roades {
		if roadGraphe[road[0]][road[1]] == 0 || roadGraphe[road[0]][road[1]] > road[2] {
			roadGraphe[road[0]][road[1]] = road[2]
			roadGraphe[road[1]][road[0]] = road[2]

		}
	}

	queue := [][2]int{}
	for villageNumber, distance := range roadGraphe[1] {
		if villageDistances[villageNumber] > distance && k >= distance {
			villageDistances[villageNumber] = distance
			queue = append(queue, [2]int{villageNumber, distance})
		}
	}

	for len(queue) != 0 {
		currentVillageNumber, currentDistance := queue[0][0], queue[0][1]
		queue = queue[1:]
		for nextVillageNumber, nextDistance := range roadGraphe[currentVillageNumber] {
			distance := currentDistance + nextDistance
			if distance <= k && villageDistances[nextVillageNumber] > distance {
				villageDistances[nextVillageNumber] = distance
				queue = append(queue, [2]int{nextVillageNumber, distance})
			}
		}
	}
	ableVillageCount := 0
	for i := range villageDistances {
		if villageDistances[i] != MAXIMUM_DISTANCE {
			ableVillageCount += 1
		}
	}
	return ableVillageCount
}
