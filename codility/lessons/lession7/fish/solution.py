def solution(A, B):
    alive_count = 0
    upstream_stack = []
    for index in range (0, len(A)):
        if B[index] == 0:
            while len(upstream_stack) != 0:
                last_upstream_fish = upstream_stack[len(upstream_stack) - 1]
                if last_upstream_fish < A[index]:
                    upstream_stack.pop()
                else:
                    break
            if len(upstream_stack) == 0:
                alive_count += 1
        else:
            upstream_stack.append(A[index])
            
            
    return alive_count + len(upstream_stack)
