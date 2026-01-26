package main

// https://leetcode.com/problems/evaluate-division/?envType=study-plan-v2&envId=leetcode-75
// @@@@@ 포기했던 문제, (a / b) * (b / c) 그래프로 접근하여 푼다는것을 생각하지 못함.
// 꼭 다시풀어볼것
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    graph := make(map[string]map[string]float64)

    for i, eq := range equations {
        u, v := eq[0], eq[1]
        val := values[i]

        if graph[u] == nil {
            graph[u] = make(map[string]float64)
        }
        if graph[v] == nil {
            graph[v] = make(map[string]float64)
        }

        graph[u][v] = val
        graph[v][u] = 1.0 / val
    }

    results := make([]float64, len(queries))

    for i, q := range queries {
        start, end := q[0], q[1]

        if graph[start] == nil || graph[end] == nil {
            results[i] = -1.0
            continue
        }

        if start == end {
            results[i] = 1.0
            continue
        }

        visited := make(map[string]bool)
        results[i] = dfs(start, end, visited, graph)
    }

    return results
}

func dfs(curr, target string, visited map[string]bool, graph map[string]map[string]float64) float64 {
    visited[curr] = true

    for neighbor, weight := range graph[curr] {
        if neighbor == target {
            return weight
        }

        if !visited[neighbor] {
            res := dfs(neighbor, target, visited, graph)
            
            if res != -1.0 {
                return weight * res
            }
        }
    }

    return -1.0
}
