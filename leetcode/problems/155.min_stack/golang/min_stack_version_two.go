package main

// other solution
type MinStack2 struct {
	stack []*entry
}

func Constructor2() MinStack2 {
	return MinStack2{
		stack: []*entry{},
	}
}

func (this *MinStack2) Push(val int) {

	minIdx := this.getMinIdx()
	if this.len() != 0 && this.stack[minIdx].val > val {
		minIdx = len(this.stack)
	}

	this.stack = append(this.stack, &entry{val, minIdx})
}

func (this *MinStack2) Pop() {
	this.stack = this.stack[:this.len()-1]
}

func (this *MinStack2) Top() int {
	return this.stack[this.len()-1].val
}

func (this *MinStack2) GetMin() int {
	return this.stack[this.getMinIdx()].val
}

func (this *MinStack2) getMinIdx() int {
	if this.len() == 0 {
		return 0
	}

	return this.stack[this.len()-1].minIdx
}

func (this *MinStack2) len() int {
	return len(this.stack)
}

type entry struct {
	val    int
	minIdx int
}
