# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution2(A):
    events = []
    for i, a in enumerate(A):
        events += [(i - a, 1), (i + a, -1)]
    events.sort(key=lambda x: (x[0], -x[1]))
    
    intersections, active_circles = 0, 0
    
    for _, event_type in events:
        intersections += active_circles * (event_type == 1)
        active_circles += event_type
        if intersections > 10_000_000:
            return -1
    
    return intersections

# Test the function with example array
A = [1, 5, 2, 1, 4, 0]
print(solution(A))  # Output should be 11

# -1 1 
# -4, 6 
# 0, 4 
# 2, 4 
# 0, 8 
# 10, 10



# -1 1 0
# 0 0 0
# -2 6 8
# 0 0 0
# 2 6 4
# 3 7 4

