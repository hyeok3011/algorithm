# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
# https://app.codility.com/programmers/lessons/6-sorting/triangle/
# 결과는 이렇게 나오면 되는듯함 P + Q > R
def solution(A):
    # Implement your solution here
    if len(A) < 2:
        return 0
    A = sorted(A)
    
    for index in range(2, len(A)):
        if A[index-2] + A[index-1] > A[index]:
            return 1
    return 0
