# https://school.programmers.co.kr/learn/courses/30/lessons/152995?language=python3
# 이 문제의 핵심은 과락의 기준을 찾는것이며 과락은 모두 걸러내며 원호보다 점수가 더 높은 사람을 찾으면 된다.
# 처음에 2중 for문을 돌며 과락을 모두 찾아낸뒤 문제를 풀었더니 시간초과가 나왔다.
# 그러다가 다른사람의 풀이 힌트를 보고 다시 풀어봤다
# 방법은 이렇다.
# 먼저 근무 태도를 기준으로 내림차순 그리고 review기존으로는 오름차순으로 정렬을 한다.
# 1번 예제를 기준으로는 [[3, 2], [3, 2], [2, 1], [1, 4]] 이렇게 나온다.
# 먼저 근무태도가 최고점인 친구에서 내려가며 계속 리뷰 점수를 갱신한다.
# 이렇게 정렬하며 계산하면 근무 태도점수는 신경쓰지않고 리뷰 점수만 신경쓰면서도 문제를 풀수있다.
# @@@@@@@@@@@@@@@@@@@@@

def solution(scores):
    target_attitude, target_review = scores[0][0], scores[0][1]
    scores = scores[1:]
    rank = 1
    scores = sorted(scores, key=lambda s: (-s[0], s[1]))
    incentive_review_threshold = scores[0][1]
    print(scores)
    
    for attitude, review in scores:
        if target_attitude < attitude and target_review < review:
            return -1

        # 여기서 과락은 필터링 된다.
        if incentive_review_threshold <= review:
            incentive_review_threshold = review
            
            if target_attitude + target_review < attitude + review:
                rank += 1

    return rank
