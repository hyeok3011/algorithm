"""
https://school.programmers.co.kr/learn/courses/30/lessons/136797
처음에 그리디 방식으로 진행하며 BFS로 했다가 틀렸다. 그 이유는
4,6
5 - 4:5, 5:6
6 - 4:6, 5:6
5 - 4:5, 4:6, 5:6
6 - 4:6, 4:6, 5:6
5 - 4:5, 5:6, 4:5, 5:6, 5:6
이런식으로 늘어나게 된다.
그래서 생각한것이 중복되는것들만 제외시키면 되겠다라는 생각으로
dict를 사용하니 풀렸다.
DP로 [][][]3중 어레이 방식으로 사용하는 경우도 있는듯한데 뭔가 순회가 많아질듯하다.
"""

def solution(numbers):
    num_position = {}
    for  i in range(1,4):
        for j in range(1,4):
            num_position[(i - 1) * 3 + j] = (i, j)
    num_position[0] = (4,2)
    
    DP = [{} for _ in range(len(numbers) + 1)]
    DP[0][(4,6)] = 0
    for i in range(0, len(numbers)):
        current_number = int(numbers[i])
        for key in DP[i]:
            left, right = key
            current_weight = DP[i][key]
            left_weight = 1 if left == current_number else get_weight(left, current_number, num_position)
            right_weight = 1 if right == current_number else get_weight(right, current_number, num_position)
            
            if current_number != right:
                if (current_number, right) not in DP[i+1]:
                    DP[i+1][(current_number, right)] = current_weight + left_weight
                elif DP[i+1][(current_number,right)] > current_weight + left_weight:
                    DP[i+1][(current_number, right)] = current_weight + left_weight
            
            if current_number != left:
                if (left, current_number) not in DP[i+1]:
                    DP[i+1][(left, current_number)] = current_weight + right_weight
                elif DP[i+1][(left,current_number)] > current_weight + right_weight:
                    DP[i+1][(left, current_number)] = current_weight + right_weight
                
    min_weight = float("inf")
    for key in DP[-1]:
        min_weight = min(min_weight, DP[-1][key])
    return min_weight


def get_weight(start:str, target:str, num_position:dict) -> int:
    start_row, start_col = num_position[start]
    target_row, target_col = num_position[target]
    diff = [abs(start_row - target_row), abs(start_col - target_col)]
    min_diff = min(diff)
    max_diff = max(diff)
    return min_diff * 3 + (max_diff - min_diff) * 2