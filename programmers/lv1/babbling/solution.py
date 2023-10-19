# https://school.programmers.co.kr/learn/courses/30/lessons/133499

# 먼저 같은 옹알이는 2번 연속으로 할수가 없어 연속되는 옹알이가 있는지 먼저 확인해본다.
# a{옹알이}ya같은 경우는 옹알이가 지워지더라도 안되기 때문에 먼저 공백으로 처리하고 나중에 
# 스페이스를 지워준다.
# 그렇게 어렵지 않은 문제인데 이상하게 10점이나 줬다 뭐지

def solution(babblings):
    answer = 0
    
    for babbling in babblings:
        for word in ["aya", "ye", "woo", "ma"]:
            if word+word in babbling:
                break
            
            babbling = babbling.replace(word, " ")
        
        if babbling.replace(" ", "") == "":
            answer += 1
    
    return answer