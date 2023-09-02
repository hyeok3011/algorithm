# https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/

def solution(N, A):
    counters = [0] * N  
    max_val = 0  
    max_counters = 0  

    for v in A:
        if v > N:
            max_counters = max_val
        else:
            counters[v-1] = max(counters[v-1], max_counters) + 1
            max_val = max(max_val, counters[v-1])

    for i in range(N):
        counters[i] = max(counters[i], max_counters)
    
    return counters