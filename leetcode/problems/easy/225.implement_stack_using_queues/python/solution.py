"""
https://leetcode.com/problems/implement-stack-using-queues/
"""

from typing import List

class MyStack:
    def __init__(self):
        self.queue: List[int] = []
        self.temp_queue: List[int] = []

    def push(self, x: int) -> None:
        self.queue.append(x)

    def pop(self) -> int:
        if len(self.queue) > 1:
            self.temp_queue = self.queue[:-1]
            self.queue = self.queue[-1:]
            
        elem: int = self.queue.pop(0)
        
        self.queue, self.temp_queue = self.temp_queue, self.queue
        return elem


    def top(self) -> int:
        if len(self.queue) > 1:
            self.temp_queue = self.queue[:-1]
            self.queue = self.queue[-1:]
            
        elem: int = self.queue.pop(0)
        self.temp_queue.append(elem)
        self.queue, self.temp_queue = self.temp_queue, self.queue
        return elem

    def empty(self) -> bool:
        if len(self.queue) == 0:
            return True
        return False
        


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
# 1
# 1 2
# 2 1 3