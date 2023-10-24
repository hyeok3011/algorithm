package main

// https://school.programmers.co.kr/learn/courses/30/lessons/148653
// 제일 작은자리수부터 0을 만들어주면 되는 문제다. 4이하면 -1씩 5이상 +1씩
// 으로풀었더니 계속 문제가 생겼다.
// 테스트 케이스가 몇개씩 통과가 안되는 문제였다.
// 아무리 생각해도 맞는거 같아서 다른 사람의 힌트를 봤다...
// 1의자리가 5일때 10의 자리를 확인해야 하는 것이였다.
// 왜냐하면 5일때는 올리던 내리던 5개의 스톤이 필요한것은 동일하지만
// 올릴때 10의자리에 영향을 주기 때문이였다. 예를들어서
// 45 일때는 내리면 5번내리고 4번 내리면 9번이지만 5번올리고 5번 내리면 10개의 스톤이 필요하게 된다.
func solution(storey int) int {
	magicStoneCount := 0
	for storey != 0 {
		reminder := storey % 10
		storey /= 10
		if reminder >= 6 || (reminder == 5 && storey%10 >= 5) {
			magicStoneCount += 10 - reminder
			storey += 1
		} else {
			magicStoneCount += reminder
		}
	}

	return magicStoneCount
}
