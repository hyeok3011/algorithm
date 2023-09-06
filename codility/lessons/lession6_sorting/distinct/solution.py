# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A):
    # Implement your solution here
    int_set = set()
    for v in A:
        int_set.add(v)
    
    return len(int_set)


