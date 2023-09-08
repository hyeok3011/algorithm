def solution(A):
    current_sum = 0
    total_sum = sum(A)
    min_diff = float('inf')

    for index in range(0 , len(A)-1):
        current_sum += A[index]
        total_sum -= A[index]
        min_diff = min(min_diff, abs(total_sum - current_sum))

    return min_diff
