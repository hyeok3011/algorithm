"""
https://school.programmers.co.kr/learn/courses/30/lessons/81302
맨해튼 거리 = (r1, c1), {r2, c2}에 각각 위치하고 있다면 그 둘의 사이의 거리는 r1-r2 + c1 - c2
단순한 구현 문제였다.
"""

def solution(places):
    answer = []
    for place in places:
        person_positions = [(r, c) for r in range(5) for c in range(5) if place[r][c] == 'P']
        if is_keep_distance(person_positions, place):
            answer.append(1)
        else:
            answer.append(0)
    return answer

def is_keep_distance(person_position, place) -> bool:
    for i in range(len(person_position)):
        r1, c1 = person_position[i]
        for j in range(i + 1, len(person_position)):
            r2, c2 = person_position[j]
            manhattan_distance = get_manhattan_distance(r1, c1, r2, c2)
            if manhattan_distance <= 2 and not has_partition_between(r1, c1, r2, c2, place):
                return False
    return True

def has_partition_between(r1, c1, r2, c2, place) -> bool:
    if r1 == r2 or c1 == c2:
        if r1 == r2 and place[r1][max(c1,c2) - 1] == "X":
            return True
        elif c1 == c2 and place[max(r1, r2) - 1][c1] == "X":
            return True
    elif place[r1][c2] == "X" and place[r2][c1] == "X":
            return True
    return False

def get_manhattan_distance(r1, c1, r2, c2) -> int:
    return abs(r1 - r2) + abs(c1 - c2)
