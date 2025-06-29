# https://school.programmers.co.kr/learn/courses/30/lessons/388352?language=python3

import itertools

def solution(n, q, anw):
    valid_code_count = 0
    
    for candidate in itertools.combinations(range(1, n + 1), 5):
        candidate_set = set(candidate)
        
        match_count = all(
            # python에서 set연산이 이렇게 가능하다니... 이걸 왜 이제알았지
            len(candidate_set & set(input_data)) == answer
            for input_data, answer in zip(q, anw)
        )
        
        if match_count:
            valid_code_count += 1
            
    return valid_code_count