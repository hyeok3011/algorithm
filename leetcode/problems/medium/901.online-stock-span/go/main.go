package main

// https://leetcode.com/problems/online-stock-span/?envType=study-plan-v2&envId=leetcode-75
// Brute Force로 풀었으나 성능 이슈로 최적화
// 뒤의 숫자가 더 작아진 경우 꼭 뒤의 숫자를 계속 유지할 필요는 없다.
type item struct {
	price int
	span  int
}

type StockSpanner struct {
	stack []item
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []item{},
	}
}

func (s *StockSpanner) Next(price int) int {
	currentSpan := 1
	idx := len(s.stack) - 1

	for idx >= 0 && s.stack[idx].price <= price {
		currentSpan += s.stack[idx].span
		idx--
	}

	s.stack = s.stack[:idx+1]

	s.stack = append(s.stack, item{
		price: price,
		span:  currentSpan,
	})

	return currentSpan
}
