package main

// https://leetcode.com/problems/remove-linked-list-elements/
// 오랜만에 풀어보는 이맛...
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	temp := ListNode{
		Val:  0,
		Next: head,
	}
	current := head
	pre := &temp
	for current != nil {
		if current.Val == val {
			pre.Next = current.Next
		} else {
			pre = current
		}
		current = current.Next
	}
	return temp.Next
}

func removeElementsRecursive(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	if head.Val == val {
		return removeElementsRecursive(head.Next, val)
	}
	head.Next = removeElementsRecursive(head.Next, val)

	return head
}
