# https://school.programmers.co.kr/learn/courses/30/lessons/340199
# 이상한 문제인게 처음부터 문제에 정답이 나와있음;;; 내가 놓친부분이 있는건가
def solution(wallet, bill):
    answer = 0
    
    while min(bill) > min(wallet) or max(bill) > max(wallet):
        if bill[0] > bill[1]:
            bill[0] = bill[0] // 2
        else:
            bill[1] = bill[1] // 2
        answer += 1
    
    return answer
