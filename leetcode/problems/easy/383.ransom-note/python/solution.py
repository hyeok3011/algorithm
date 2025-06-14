# https://leetcode.com/problems/ransom-note/description/
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        char_count = {}
        
        for char in magazine:
            char_count[char] = char_count.get(char, 0) + 1
        
        for char in ransomNote:
            if char not in char_count or char_count[char] <= 0:
                return False
            char_count[char] -= 1
        
        return True

# 지렸다
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        
        for i in ransomNote:
            if i in magazine:
                magazine = magazine.replace(i,"",1)
            else: return False
        
        return True