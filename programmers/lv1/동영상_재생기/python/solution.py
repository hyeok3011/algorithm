# https://school.programmers.co.kr/learn/courses/30/lessons/340213
# 항상 쉬운듯 어려운 프로그래머스
def solution(video_len, pos, op_start, op_end, commands):
    pos = str_time_to_second(pos)
    op_start = str_time_to_second(op_start)
    op_end = str_time_to_second(op_end)
    video_len = str_time_to_second(video_len)
    
    for command in commands:
        if is_opening_pos(pos, op_start, op_end):
            pos = op_end
        if command == "prev":
            pos -= 10
            pos = max(0, pos)
        else:
            pos += 10
            pos = min(pos, video_len)
        if is_opening_pos(pos, op_start, op_end):
            pos = op_end
                    
    return str(pos // 60).zfill(2) + ":" + str(pos % 60).zfill(2)

def is_opening_pos(pos, op_start, op_end):
    return op_start <= pos <= op_end

def str_time_to_second(pos):
    minute, second = pos.split(":")
    return int(minute) * 60 + int(second)