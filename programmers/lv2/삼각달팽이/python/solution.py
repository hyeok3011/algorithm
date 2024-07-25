"""
https://school.programmers.co.kr/learn/courses/30/lessons/68645
"""
from typing import List

def solution(n: int) -> List[int]:
    box_count: int = (n * (n + 1)) // 2

    triangle: List[List[int]] = []
    for i in range(n):
        triangle.append([0] * (i + 1))

    row: int = 0
    col: int = 0
    row_step: int = 1
    col_step: int = 0

    for i in range(1, box_count + 1):
        triangle[row][col] = i

        next_row = row + row_step
        next_col = col + col_step

        if next_row >= n or next_col > next_row or triangle[next_row][next_col] != 0:
            if row_step == 1 and col_step == 0:
                row_step = 0
                col_step = 1
            elif row_step == 0 and col_step == 1:
                row_step = -1
                col_step = -1
            elif row_step == -1 and col_step == -1:
                row_step = 1
                col_step = 0

        row += row_step
        col += col_step

    answer: List[int] = []
    for l in triangle:
        answer += l

    return answer

print(solution(10))
