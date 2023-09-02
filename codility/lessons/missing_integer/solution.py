# https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/

def solution(A):
    # Implement your solution here
    A = [item for item in A if item >= 1]
    A = sorted(A)

    if len(A) == 0:
        return 1

    if A[0] != 1:
        return 1

    pre_val = 1
    for index in range(1, len(A)):
        if pre_val != A[index] and pre_val + 1 != A[index]:
            return pre_val + 1
        pre_val = A[index]

    return pre_val + 1