# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")


"""
https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_slice_sum/start/
어디서 많이 본듯한 서브 어레이합중 가장 큰 합을 찾는 문제이다.
"""

def solution(A):
    if len(A) == 0:
        return 0
    
    max_sum = A[0]
    current_sum = A[0]
    for i in range (1, len(A)):
        if current_sum < 0:
            current_sum = 0
            
        current_sum += A[i]
        max_sum = max(max_sum, current_sum)

            
    return max_sum
