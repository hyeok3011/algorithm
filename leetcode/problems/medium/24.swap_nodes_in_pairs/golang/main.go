package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/swap-nodes-in-pairs/description/
// Runtime 2 ms Beats 67.22% Memory 2 MB Beats 50.37%

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var current, next, pre *ListNode
	current = head
	head = head.Next
	next = current.Next
	for {
		current.Next = next.Next
		next.Next = current

		if pre != nil {
			pre.Next = next
		}

		pre = current

		current = current.Next
		if current == nil || current.Next == nil {
			break
		}
		next = current.Next
	}

	return head
}

func otherPersonSolution(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	head.Next = swapPairs(head.Next.Next)
	result.Next = head
	return result
}
