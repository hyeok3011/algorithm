package main

type MyQueue struct {
	outStack *Stack
	inStack  *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		outStack: NewStack(),
		inStack:  NewStack(),
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.outStack.IsEmpty() {
		q.transfer()
	}

	return q.outStack.Pop()
}

func (q *MyQueue) transfer() {
	for q.inStack.Size() > 0 {
		q.outStack.Push(q.inStack.Pop())
	}
}

func (q *MyQueue) Peek() int {
	if q.outStack.IsEmpty() {
		q.transfer()
	}
	return q.outStack.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.inStack.IsEmpty() && q.outStack.IsEmpty()
}

type Stack struct {
	elem []int
}

func NewStack() *Stack {
	return &Stack{
		elem: []int{},
	}
}

func (s *Stack) Push(v int) {
	s.elem = append(s.elem, v)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.elem) - 1
	v := s.elem[lastIndex]
	s.elem = s.elem[:lastIndex]
	return v
}

func (s *Stack) Peek() int {
	return s.elem[len(s.elem)-1]
}
func (s *Stack) Size() int {
	return len(s.elem)
}
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
