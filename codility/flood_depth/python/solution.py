# 이 전체 지역이 침수된 후 최대 수심을 알고 싶습니다.
# 물이 잠길수 없는 환경에서는 0을 리턴
# 아니면 최대 수심

MAX_WALL_HEIGHT = 100000000
def solution(A):
    # Implement your solution here
    maximum = 0

    left_wall = 0
    low_wall = MAX_WALL_HEIGHT
    for height in A:
        maximum = max(maximum, min(left_wall, height) - low_wall)
        if low_wall > height:
            low_wall = height

        if height >= left_wall:
            left_wall = height
            low_wall = MAX_WALL_HEIGHT

    return maximum

# two pointer algorithm
def solution2(A):
    N = len(A)
    if N < 3:  
        return 0
    
    left, right = 0, N - 1
    max_depth = 0
    left_wall, right_wall = A[left], A[right]
    
    while left <= right:
        if left_wall <= right_wall:
            left_wall = max(left_wall, A[left])
            max_depth = max(max_depth, left_wall - A[left])
            left += 1
        else:
            right_wall = max(right_wall, A[right])
            max_depth = max(max_depth, right_wall - A[right])
            right -= 1
            
    return max_depth
