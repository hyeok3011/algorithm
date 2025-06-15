# https://school.programmers.co.kr/learn/courses/30/lessons/388351
def solution(schedules, timelogs, startday):
    answer = 0
    for i, scheduler in enumerate(schedules):
        current_day = startday
        accept_time = scheduler + 10
        if accept_time % 100 >= 60:
            accept_time += 40
            
        valid = True
        for time_log in timelogs[i]:
            if current_day < 6:
                if accept_time < time_log:
                    valid = False
                    break
            current_day = current_day % 7 + 1

        if valid: answer += 1
            
        
    return answer

