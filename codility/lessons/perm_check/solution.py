# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
# https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/
def solution(A):
    # Implement your solution here
    A = sorted(A)
    for index in range(0, len(A)):
        if A[index] != index+1:
            return 0
    
    return 1

def solution2(A):
    elem_set = set(A)
    max_val = max(A)
        
    if max_val == len(elem_set) and len(elem_set) == len(A):
        return 1
    
    return 0