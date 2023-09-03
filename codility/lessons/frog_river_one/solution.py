# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(X, A):
    # Implement your solution here
    leaf_position_set = set()
    min_time = -1
    for time in range(0, len(A)):
        if A[time] <= X:
            leaf_position_set.add(A[time])
        
        if len(leaf_position_set) == X:
            min_time = time
            break
    
    return min_time
