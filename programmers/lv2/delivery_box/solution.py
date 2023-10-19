# https://school.programmers.co.kr/learn/courses/30/lessons/131704
# 단순한 stack 문제이다. 조기 종료를 위해 for문마다 조기 종료 조건을 추가했다.
def solution(orders):
    max_order = max(orders)
    stack = []
    order_index = 0

    for i in range(1, max_order + 1):
        stack.append(i)
        while stack and orders[order_index] == stack[-1]:
            order_index += 1
            stack.pop()
        
        if order_index == len(orders) or orders[order_index] < i:
            break
        
    return order_index

solution([4, 3, 1, 2, 5, 6]) == 2