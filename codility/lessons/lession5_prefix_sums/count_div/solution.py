def solution(A, B, K):
    # Implement your solution here
    b_div_count = B // K
    
    a_div_count = 0
    if A != 0:
        a_div_count = A // K
    
    div_count = b_div_count - a_div_count 
    if A % K == 0:
        div_count += 1
    
    return div_count