# https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/
def solution(X, Y, D):
    # Implement your solution here
    if X == Y:
        return 0

    jump_count = int((Y - X ) / D)
    remainder = (Y - X ) % D
    if remainder > 0:
        jump_count += 1
        
    return jump_count


import math
def solution2(X, Y, D):
    return math.ceil((Y - X) / D)