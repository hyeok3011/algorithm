from collections import deque, defaultdict

# https://school.programmers.co.kr/learn/courses/30/lessons/132266?language=python3
# 먼저 현재 위치에서 목표 지점까지 최단거리를 구해야 하기 때문에 거리 찾는 부분에서의 알고리즘은 BFS를 써야겠다 생각했다.
# 여기서 풀수있는 방법은 2가지 인거같다.
# 1. souces를 순회하며 destination를 찾는 방법 (너무 무식하기 때문에 메모이제이션은 기본적으로 같이 사용)
# 2. 다른 방법은 graph를 순회하며 각 지점별로 최단 거리를 미리 구해놓는 방법
# n의 크기가 큰 상태에서 데이터가 극단적으로 source에 몰려있지 않다면 분명 두번째 방법이 더 빠를거라고 판단하여
# 두번째 방법으로 풀었다.
def solution(n, roads, sources, destination):
    graph = defaultdict(list)

    for road in roads:
        graph[road[0]].append(road[1])
        graph[road[1]].append(road[0])

    distances = {destination: 0}
    queue = deque([destination])
    
    while queue:
        current_position = queue.popleft()
        for next_position in graph[current_position]:
            if next_position not in distances:
                distances[next_position] = distances[current_position] + 1
                queue.append(next_position)

    answer = []
    for source in sources:
        answer.append(distances.get(source, -1))

    return answer

