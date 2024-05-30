function solution(N, road, K) {
    //    [[1,2,1],[2,3,3],[5,2,2],[1,4,2],[5,3,1],[5,4,2]]
        let graph = {};
        road.forEach(r => {
            let [start, end, distance] = [r[0], r[1], r[2]];
            graph[start] = graph[start] || [];
            graph[start].push([end, distance]);
            graph[end] = graph[end] || [];
            graph[end].push([start, distance]);
        })

        let queue = [[1, 0]];    
        let campDistance = new Array(N + 1).fill(K + 1);
        campDistance[1] = 0;
        while (queue.length > 0) {
            let queLen = queue.length;
            let currentNode = queue.shift();
            let [start, currentDistance] = [currentNode[0], currentNode[1]];
            
            if (currentDistance > K || currentDistance > campDistance[start]) {
                continue;
            }
            
            for (let i = 0; i < queLen; i++) {
                graph[start].forEach(node => {
                    if (campDistance[node[0]] > node[1] + currentDistance) {
                        campDistance[node[0]] = node[1] + currentDistance
                        queue.push([node[0], node[1] + currentDistance])
                    }
                })
            }
        }
        console.log(campDistance)
        return campDistance.filter(n => n <= K).length;
    }
