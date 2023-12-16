from typing import List


"""
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
간단하게 슬라이스에서 받는 스톡 가격에서 모든 이득을 취하면 되는 문제...
"""

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        current_stock_price = prices[0]
        total_profit = 0
        for price in prices:
            if current_stock_price < price:
                total_profit += price - current_stock_price
                current_stock_price = price
            elif current_stock_price > price:
                current_stock_price = price
                
        return total_profit