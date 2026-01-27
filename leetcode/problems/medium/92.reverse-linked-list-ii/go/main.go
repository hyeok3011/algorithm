package main

//https://leetcode.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	nodes := []*ListNode{&ListNode{Val: 0, Next: nil}}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	nodes = append(nodes, nil)

	for i := 0; i <= (right-left)/2; i++ {
		nodes[left+i], nodes[right-i] = nodes[right-i], nodes[left+i]
	}

	for i := left - 1; i <= right; i++ {
		nodes[i].Next = nodes[i+1]
	}

	return nodes[1]
}

// 힌트보고 다시 풀어본 문제. 간단하지만 쉽지 않았음. @@@@@
func anotherSolution(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{0, head}

	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	curr := prev.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
