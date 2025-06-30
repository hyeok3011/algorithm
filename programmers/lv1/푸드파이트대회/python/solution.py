# https://school.programmers.co.kr/learn/courses/30/lessons/134240
def solution(food):
    left_player = ""
    right_player = ""
    for i in range(1, len(food)):
        left_player += str(i) * int(food[i] // 2)
        right_player = str(i) * int(food[i] // 2) + right_player
    
    return left_player + "0" + right_player