# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
# https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/
# 0 ≤ P < Q < N인 정수 쌍(P, Q)을 배열 A의 슬라이스라고 합니다.
# 슬라이스(P, Q)의 평균은 A[P] + A[P + 1] + ... + A[Q]의 합을 슬라이스 길이로 나눈 값입니다.
# 정확하게 말하면 평균은 (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1)과 같습니다.
def solution(A):
    min_avg = 10000
    min_avg_index = 0
    
    for i in range(len(A) - 2):
        front_two_avg = (A[i] + A[i + 1]) / 2
        front_three_avg = (A[i] + A[i + 1] + A[i + 2]) / 3
        
        if front_two_avg < min_avg:
            min_avg = front_two_avg
            min_avg_index = i
        
        if front_three_avg < min_avg:
            min_avg = front_three_avg
            min_avg_index = i
    
    if (A[len(A) - 1] + A[len(A) - 2]) / 2 < min_avg:
        min_avg = (A[len(A) - 1] + A[len(A) - 2]) / 2
        min_avg_index = len(A) - 2
    
    return min_avg_index
            

        
        
# 10, 9, 8, 7, 6, 5, 4, 3, 2 ,1
#  19  17 15 13 11 9 7 5 3 
solution([2, 1, 2, 2, 5, 1, 2, 1])
# 2, 1, 2, 2, 5, 1, 2, 1
# 3 3 4 7 6 3 3

# -10, 1, 5, -5, -7, -1, -5, -7, -10
# -9, 6, 0, -12, -8, -6, -12, -17

