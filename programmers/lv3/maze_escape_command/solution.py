# https://github.com/Juniork725/coding_test/blob/main/%EB%82%9C%EC%9D%B4%EB%8F%843/%EB%AF%B8%EB%A1%9C%20%ED%83%88%EC%B6%9C%20%EB%AA%85%EB%A0%B9%EC%96%B4.md
# 다른사람의 파이썬 풀이법이다.
# 복잡한 문제를 이렇게 단순화할수 있다니....
# @@@@@@@@@@@2
def solution(n, m, x, y, r, c, k):
    answer = ''
    dist = abs(x-r)+abs(y-c)
    k -= dist
    if k < 0 or k%2 != 0:
        return "impossible"
    
    direction = {'d':0, 'l':0, 'r':0, 'u':0}
    if x > r:
        direction['u'] += x-r
    else:
        direction['d'] += r-x
    if y > c:
        direction['l'] += y-c
    else:
        direction['r'] += c-y
        
    answer += 'd'*direction['d']
    d = min(int(k/2), n-(x+direction['d']))
    answer += 'd'*d
    direction['u'] += d
    k -= 2*d
    
    answer += 'l'*direction['l']
    l = min(int(k/2), y-direction['l']-1)
    answer += 'l'*l
    direction['r'] += l
    k -= 2*l
    
    answer += 'rl'*int(k/2)
    answer += 'r'*direction['r']
    answer += 'u'*direction['u']
    return answer