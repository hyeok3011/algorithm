package main

// https://leetcode.com/problems/min-stack/description/
// Runtime 13 ms Beats 92.30% Memory 7.1 MB Beats 94.16%
type MinStack struct {
	min      int
	elements []int
}

func Constructor() MinStack {
	return MinStack{min: 2147483647}
}

func (this *MinStack) Push(val int) {
	if this.min > val {
		this.min = val
	}
	this.elements = append(this.elements, val)
}

func (this *MinStack) Pop() {
	topElement := this.elements[len(this.elements)-1]
	this.elements = this.elements[:len(this.elements)-1]

	if topElement == this.min {
		this.min = 2147483647
		for _, val := range this.elements {
			if this.min > val {
				this.min = val
			}
		}
	}
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
