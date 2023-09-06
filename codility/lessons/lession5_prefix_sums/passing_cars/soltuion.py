# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
# P가 동쪽으로 이동하고 Q가 서쪽으로 이동할 때 0 ≤ P < Q < N인 한 쌍의 자동차(P, Q)가 지나가고 있다고 말합니다.
def solution(A):
    # Implement your solution here
    east_count = 0
    west_count = 0
    
    for index, v in enumerate(A):
        if v == 0:
            A[index] = west_count * 10
            east_count += 1
        else:
            west_count += 1
    
    pair_count = 0
    for v in A:
        if v != 1:
            pair_count += west_count - (v / 10)
    

    if 1000000000 < pair_count:
        return -1
    return int(pair_count)

def solution2(A):
    east_count = 0  
    passing_cars = 0  

    for car in A:
        if car == 0:
            east_count += 1  
        else:
            passing_cars += east_count  

        
        if passing_cars > 1000000000:
            return -1

    return passing_cars