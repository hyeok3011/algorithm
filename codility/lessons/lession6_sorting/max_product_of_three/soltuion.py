# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A):
    # Implement your solution here
    A = sorted(A)
    
    first_two_liast_mul_val = A[0] * A[1] * A[-1]
    last_three_mul_val = A[-1] * A[-2] * A[-3]
    
    if first_two_liast_mul_val > last_three_mul_val:
        return first_two_liast_mul_val
    
    return last_three_mul_val
