"""
https://school.programmers.co.kr/learn/courses/30/lessons/142085
처음에는 완전 그리디 방식으로 진행했다가 틀리는 케이스가 있어 heap을 사용하였다.
병력이 남아 있는경우에는 먼저 남아있는 병력으로 무조건 디펜스를 하다가 병력이 부족해질 경우
무적권을 사용을 하는데 최대한의 효율을 위해서 지금까지 막아온 케이스중 가장 큰 케이스를 사용하도록 하였다.
"""
import heapq

def solution(n, k, enemy):
    if k >= len(enemy):
        return len(enemy)
    
    defensed_enemy = []
    index = 0
    heapq.heapify(defensed_enemy)
    while index < len(enemy):
        if n >= enemy[index]:
            n -= enemy[index]
            heapq.heappush(defensed_enemy, -enemy[index])
        elif k > 0:
            if defensed_enemy and -defensed_enemy[0] > enemy[index]:
                n += -heapq.heappop(defensed_enemy)
                n -= enemy[index]
                heapq.heappush(defensed_enemy, -enemy[index])
            k -= 1
        else:
            break
        index += 1

    return index