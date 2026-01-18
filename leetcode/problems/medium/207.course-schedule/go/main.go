package main

// https://leetcode.com/problems/course-schedule/?envType=study-plan-v2&envId=leetcode-75
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, req := range prerequisites {
		graph[req[1]] = append(graph[req[1]], req[0])
	}

	state := make([]int, numCourses)

	var hasCycle func(curr int) bool
	hasCycle = func(curr int) bool {
		if state[curr] == 1 {
			return true
		}
		if state[curr] == 2 {
			return false
		}

		state[curr] = 1

		for _, neighbor := range graph[curr] {
			if hasCycle(neighbor) {
				return true
			}
		}

		state[curr] = 2
		return false
	}

	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}

	return true
}

// Kahns방식으로 푼 다른 사람의 풀이
func canFinishKahs(numCourses int, prerequisites [][]int) bool {
	adjList := make([][]int, numCourses)
	inDeg := make([]int, numCourses)

	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		adjList[b] = append(adjList[b], a)
		inDeg[a]++
	}

	q := []int{}
	for i, d := range inDeg {
		if d == 0 {
			q = append(q, i)
		}
	}

	processedCount := 0

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		processedCount++

		for _, neigh := range adjList[node] {
			inDeg[neigh]--
			if inDeg[neigh] == 0 {
				q = append(q, neigh)
			}
		}
	}

	return processedCount == numCourses

}
