# https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/

# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(S: str):
    # Implement your solution here
    if S.isspace():
        return 1
    bracket_pair = {
        "}": "{",
        "]": "[",
        ")": "(",
    }
    stack = []

    for v in S:
        if  v in ["{", "[", "("]:
            stack.append(v)
        elif v in ["}" , "]" , ")"]:
            if len(stack) == 0 or bracket_pair[v] != stack.pop():
                return 0
    
    if len(stack) != 0:
        return 0
            
    return 1
