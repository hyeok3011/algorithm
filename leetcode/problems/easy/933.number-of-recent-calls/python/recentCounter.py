# https://leetcode.com/problems/number-of-recent-calls/?envType=study-plan-v2&envId=leetcode-75

from collections import deque

class RecentCounter:
    TIME_WINDOW_MS = 3000
    def __init__(self):
        self.request_queue = deque()

    def ping(self, t: int) -> int:
        self.request_queue.append(t)
        self._prune_old_requests(t)
        return len(self.request_queue)
        
    def _prune_old_requests(self, t: int) -> None:
        cutoff_time: int = t - self.TIME_WINDOW_MS
        while self.request_queue[0] < cutoff_time:
            self.request_queue.popleft()
        
# deque를 사용하지 않는 코드
class RecentCounter2:
    TIME_WINDOW_MS = 3000
    def __init__(self):
        self.request_queue = []
        self.cutoff_index = 0

    def ping(self, t: int) -> int:
        self.request_queue.append(t)

        cutoff_time: int = t - self.TIME_WINDOW_MS
        while self.request_queue[self.cutoff_index] < cutoff_time:
            self.cutoff_index += 1

        return len(self.request_queue) - self.cutoff_index