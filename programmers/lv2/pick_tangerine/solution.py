from collections import Counter

# https://school.programmers.co.kr/learn/courses/30/lessons/138476
# 생각보다 lv2문제치고 너무 간단했다. 
# size별로 몇개의 귤이 있는지 내림차순으로 정렬하고 k이상 될때까지 반복해주면 된다.
def solution(k, tangerine):
    answer = 0
    tangerineSizes = Counter(tangerine)
    
    for count in sorted(tangerineSizes.values(), reverse=True):
        k -= count
        answer += 1
        
        if k <= 0:
            break
    
    return answer