# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(A):
    element_hash_table = {}
    for i, v in enumerate(A):
        if v in element_hash_table:
            element_hash_table[v] += 1
        else:
            element_hash_table[v] = 1
        
        if element_hash_table[v] > len(A) / 2:
            return i
        
    
    return -1