package main

// parents? root를 한번만 찾으면 된다.
func validPath(n int, edges [][]int, source int, destination int) bool {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}

		return parents[x]
	}

	union := func(src, desc int) {
		rootSrc, rootDesc := find(src), find(desc)
		if rootSrc != rootDesc {
			parents[rootSrc] = rootDesc
		}
	}

	for _, edge := range edges {
		start, end := edge[0], edge[1]
		union(start, end)
	}

	return find(source) == find(destination)
}

func validPathDFS(n int, edges [][]int, source int, destination int) bool {
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)
	var dfs func(int) bool
	dfs = func(current int) bool {
		if current == destination {
			return true
		}
		visited[current] = true
		for _, next := range graph[current] {
			if !visited[next] && dfs(next) {
				return true
			}
		}
		return false
	}

	return dfs(source)
}
