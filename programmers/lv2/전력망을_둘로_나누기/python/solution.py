"""
https://school.programmers.co.kr/learn/courses/30/lessons/86971
다른사람들중에 union find방식으로 푼사람도 있는데 좋은 방법인듯 하다.
"""

def solution(n, wires):
    graph = {}
    for wire in wires:
        if wire[0] not in graph:
            graph[wire[0]] = {}
        if wire[1] not in graph:
            graph[wire[1]] = {}
        graph[wire[0]][wire[1]] = True
        graph[wire[1]][wire[0]] = True
    
    def dfs(start, tower_count):
        for end in graph[start]:
            if not graph[start][end]:
                continue
            graph[start][end] = False
            graph[end][start] = False
            tower_count = dfs(end, tower_count + 1)
            graph[start][end] = True
            graph[end][start] = True
            
        return tower_count

    min_diff = 100
    for wire in wires:
        start, end = wire
        graph[start][end] = False
        graph[end][start] = False
        tower_count = dfs(end, 1)
        graph[start][end] = True
        graph[end][start] = True
        min_diff = min(min_diff, abs(n - (tower_count * 2)))
    return min_diff

print(solution(9, [[1,3],[2,3],[3,4],[4,5],[4,6],[4,7],[7,8],[7,9]]))