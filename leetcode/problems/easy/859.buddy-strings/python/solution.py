# https://leetcode.com/problems/buddy-strings/description/
# 위치가 2개의 위치를 바꿨을때 s 와 goal이 동일하다면 True 아니면 False
# 서로 같은 문자인경우 어쨋든 바꿀수 있으므로 True

# 만약 같은데 겹치는 원소가 있으면 바로 true
# 같지 않은경우 위치가 다른 원소가 2개있고 바꿨을때 같으면 true
class Solution:
    def buddyStrings(self, s: str, goal: str) -> bool:
        if len(s) != len(goal):
            return False
        
        if s == goal:
            return len(set(s)) < len(s)
        
        diff_position = []
        for i in range(len(s)):
            if s[i] != goal[i]:
                diff_position.append(i)
        
        if len(diff_position) != 2:
            return False
        
        i, j = diff_position
        return s[i] == goal[j] and s[j] == goal[i]