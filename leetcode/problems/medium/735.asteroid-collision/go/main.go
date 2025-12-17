package main

// https://leetcode.com/problems/asteroid-collision/?envType=study-plan-v2&envId=leetcode-75
func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, ast := range asteroids {
		alive := true
		for alive && ast < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top < -ast {
				stack = stack[:len(stack)-1]
				continue
			} else if top == -ast {
				stack = stack[:len(stack)-1]
				alive = false
			} else {
				alive = false
			}
		}
		if alive {
			stack = append(stack, ast)
		}
	}
	return stack
}
