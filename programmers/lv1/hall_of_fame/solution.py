# https://school.programmers.co.kr/learn/courses/30/lessons/138477

def solution(k, scores):
    answer = []
    
    queue = []
    for score in scores:
        queue.append(score)
        if len(queue) > k:
            queue.remove(min(queue))            

        answer.append(min(queue))


    return answer
    