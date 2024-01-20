"""
https://leetcode.com/problems/reverse-words-in-a-string/
간단하게 들어오는 문자열을 단어 기준으로 뒤집으며 앞뒤 공백을 제거하는 문제이다.
다른사람들의 풀이를 보니 정말 쉽게 풀었다;;
"""


class Solution:
    def reverseWords(self, s: str) -> str:
        reversedWord = ""
        i = len(s) - 1

        while i >= 0:
            if s[i] == " ":
                i -= 1
                continue
            
            j = i
            while i >= 0 and s[i] != " ":
                i -= 1
            
            reversedWord += s[i+1:j+1] + " "
        
        return reversedWord.rstrip()
    
    def otherSolution(self, s: str) -> str:
        x = s.split()
        return " ".join(x[::-1])
    
