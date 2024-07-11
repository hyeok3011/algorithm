/*
https://school.programmers.co.kr/learn/courses/30/lessons/118669
풀긴 풀었는데 다익스트라 알고리즘을 사용하여 노드별 가장 작은 Intensity를 구한다면 더 빠르게 풀리지 않을까 싶다.
*/ 
function solution(n, paths, gates, summits) {
    const graph = {};

    for (const path of paths) {
        const [start, end, intensity] = path;
        if (!graph[start]) {
            graph[start] = [];
        }
        if (!graph[end]) {
            graph[end] = [];
        }

        graph[start].push([end, intensity]);
        graph[end].push([start, intensity]);
    }

    const visited = {};
    for (let i = 1; i <= n; i++) {
        if (graph[i]) {
            graph[i].sort((a, b) => a[1] - b[1]);
        }
        visited[i] = Infinity;
    }

    const summitSet = new Set(summits);
    const gatesSet = new Set(gates);
    let maxIntensity = Infinity;
    let minSummit = Infinity;
    summits.sort((a, b) => a - b);
    
    for (const start of summits) {
        const queue = [];
        visited[start] = 0;
        for (const v of graph[start]) {
            queue.push([v[0], v[1], v[1]]);
        }

        while (queue.length > 0) {
            const len = queue.length;
            for (let i = 0; i < len; i++) {
                let [point, pointIntensity, pathMaxIntensity] = queue.shift();

                pathMaxIntensity = Math.max(pointIntensity, pathMaxIntensity);
                if (pathMaxIntensity >= visited[point] || pathMaxIntensity >= maxIntensity) {
                    continue
                }

                visited[point] = pathMaxIntensity;

                if (gatesSet.has(point)) {
                    if (pathMaxIntensity < maxIntensity) {
                        maxIntensity = pathMaxIntensity;
                        minSummit = start;
                    }
                    continue;
                }
                
                for (const v of graph[point]) {
                    if (visited[v[0]] <= pathMaxIntensity || visited[v[0]] <= v[1] || summitSet.has(v[0])) {
                        continue;
                    }
                    queue.push([v[0], v[1], pathMaxIntensity]);
                }
            }
        }
    }
    
    return [minSummit, maxIntensity];
}


console.log(solution(
    6,
    [[1, 2, 3], [2, 3, 5], [2, 4, 2], [2, 5, 4], [3, 4, 4], [4, 5, 3], [4, 6, 1], [5, 6, 1]],
    [1, 3],
    [5]
))