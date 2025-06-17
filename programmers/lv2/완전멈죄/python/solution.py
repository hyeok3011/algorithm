# https://school.programmers.co.kr/learn/courses/30/lessons/389480
def solution(infos, n, m):
    infos.sort(key=lambda x: x[0] / x[1], reverse=True)
    a = n
    b = m
    for info in infos:
        if 0 < b - info[1]:
            b -= info[1]
        elif 0 < a - info[0]:
            a -= info[0]
        else:
            return -1
    
    return n - a