"""
https://school.programmers.co.kr/learn/courses/30/lessons/64065?language=javascript
"""
from typing import List

def solution(s: str):
    s = s.replace("{{", "").replace("}}", "")
    tuples: List[List[str]] = []
    for tuple in s.split("},{"):
        tuples.append(tuple.split(","))
    
    tuples.sort(key=len)
    
    answer: List[int] = []
    num_set: set = set()
    for elements in tuples:
        for elem in elements:
            if elem not in num_set:
                num_set.add(elem)
                answer.append(int(elem))
        
    return answer