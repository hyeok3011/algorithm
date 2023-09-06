
def solution(A):
    # Implement your solution here
    element_hash_table = {}
    dominator = 0
    dominator_count = 0
    for v in A:
        if v in element_hash_table:
            element_hash_table[v] += 1
        else:
            element_hash_table[v] = 1
        
        if element_hash_table[v] > len(A) / 2:
            dominator = v
            dominator_count = element_hash_table[dominator]
            
    
    dominator_occurs = [0] * len(A)
    cnt = 0
    for i, v in enumerate(A):
        if v == dominator:
            cnt += 1
            dominator_occurs[i] = cnt
        else:
            dominator_occurs[i] = cnt
    
    equi_leader_count = 0
    for i in range(0, len(A)):
        if ((i+1) / 2 < dominator_occurs[i]) and (len(A) - (i+1)) /2 < dominator_count -dominator_occurs[i]:
            equi_leader_count += 1
            
    return equi_leader_count
            

