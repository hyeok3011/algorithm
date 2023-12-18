"""
https://leetcode.com/problems/maximum-gap/description/
정렬되지 않은 정수 슬라이스 중 정렬 기준 nums[i+1] - nums[i]의 차가 가장 큰 값을 구하는 문제이다.
처음에는 간단하게 정렬 후 loop를 돌면 되겠다 싶었지만 문제를 다시보고다니 linear한 시간을
가져야 한다고 나와있는데 sort를 쓰지 말라는 뜻으로 이해하면 될듯하다...ㅠㅠ
python의 default sort알고리즘은 tim sort으로 시간복잡도는 O(n Log n)이다.
https://en.wikipedia.org/wiki/Bucket_sort#
bucket sort알고리즘을 써야할듯 하다.
분명 내가 bucket sort을 쓰더라도 python의 default sort알고리즘 보다 더 느릴것다고 예상했는데
마법같이 진짜 더 느리다.
버킷 배열 사용 및 할당 오버헤드로 예상은 되는데 다음에 확인해보자...
"""

class Solution:
    def maximumGap(self, nums):
        if len(nums) < 2:
            return 0

        min_val = nums[0]
        max_val = nums[0]
        for v in nums:
            if v > max_val:
                max_val = v
            elif v < min_val:
                min_val = v
            
        if min_val == max_val:
            return 0

        bucket_size = max(1, (max_val - min_val) // (len(nums) - 1))
        bucket_count = (max_val - min_val) // bucket_size + 1

        minBucket = [float('inf')] * bucket_count
        maxBucket = [float('-inf')] * bucket_count

        for num in nums:
            idx = (num - min_val) // bucket_size
            minBucket[idx] = min(minBucket[idx], num)
            maxBucket[idx] = max(maxBucket[idx], num)

        max_gap, prev_max = 0, min_val
        for i in range(bucket_count):
            if minBucket[i] == float('inf'):
                continue
            max_gap = max(max_gap, minBucket[i] - prev_max)
            prev_max = maxBucket[i]

        return max_gap
