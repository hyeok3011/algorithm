# https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/
def solution(A):
    # Implement your solution here
    xor_sum = 0
    for v in range(0 , len(A)+1):
        xor_sum ^= v+1

    for v in A:
        xor_sum ^= v
        
    return xor_sum

def solution2(A):
    if len(A) == 0:
        return 1
    
    A = sorted(A)
    for index in range(0, len(A)):
        if A[index] != index+1:
            return index + 1
    
    return len(A) + 1
    
