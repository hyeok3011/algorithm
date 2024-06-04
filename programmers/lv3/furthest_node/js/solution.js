// https://school.programmers.co.kr/learn/courses/30/lessons/49189
// 가장 멀리 떨어진 노드란 최단경로로 이동했을 때 간선의 개수가 가장 많은 노드들을 의미
function solution(n, edge) {
    let graph = {}
    edge.forEach(edge => {
        let [start, end] = [edge[0], edge[1]];
        graph[start] = graph[start] || [];
        graph[start].push(end);
        graph[end] = graph[end] || [];
        graph[end].push(start);
    })
    
    let visited = new Set([1, ...graph[1]]);

    let queue = [...graph[1]];
    let furthestEdgeCount = 0;
    while (queue.length > 0) {
        furthestEdgeCount = queue.length;
        for (let i = 0; i < furthestEdgeCount; i++) {
            let currentNode = queue.shift();
            graph[currentNode].forEach(node => {
                if (!visited.has(node)) {
                    visited.add(node)
                    queue.push(node);
                }
            })
        }
    }
    
    return furthestEdgeCount;
}

console.log(solution(6, [[3, 6], [4, 3], [3, 2], [1, 3], [1, 2], [2, 4], [5, 2]]))