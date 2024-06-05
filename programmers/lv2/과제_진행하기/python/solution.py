#  https://school.programmers.co.kr/learn/courses/30/lessons/176962

def solution(plans):
    completed_tasks = []
    paused_tasks = []
    plans = sorted(plans, key=lambda plan: plan[1])
    
    for i in range(0, len(plans) - 1):
        task_name = plans[i][0]
        
        time_diff = calculate_time_diff(plans[i][1], plans[i+1][1])
        remaining_time = time_diff - int(plans[i][2])
        if (remaining_time < 0):
            paused_tasks.append([task_name, -remaining_time])
        else:
            completed_tasks.append(task_name)
            while remaining_time > 0 and len(paused_tasks) > 0:
                remaining_time -= paused_tasks[-1][1]
                if remaining_time >= 0:
                    completed_tasks.append(paused_tasks[-1][0])
                    paused_tasks.pop()
                else:
                    paused_tasks[-1][1] = -remaining_time

    completed_tasks.append(plans[-1][0])
                    
    for i in range(len(paused_tasks)-1, -1, -1):
        completed_tasks.append(paused_tasks[i][0])
    return completed_tasks

def calculate_time_diff(a, b):
    a_hour, a_minute = map(int, a.split(":"))
    b_hour, b_minute = map(int, b.split(":"))
    diff = (b_hour - a_hour) * 60 + (b_minute - a_minute)
    return diff
