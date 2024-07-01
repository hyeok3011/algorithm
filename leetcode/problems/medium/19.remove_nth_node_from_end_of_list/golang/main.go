package main

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

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
// Runtime 2 ms Beats 50.59% Memory 2.3 MB Beats 15.97%
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	nodes := []*ListNode{}
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	current = nodes[len(nodes)-n]
	if len(nodes)-n == 0 {
		if len(nodes) == 1 {
			return nil
		}
		head = nodes[1]
	} else {
		pre := nodes[len(nodes)-n-1]
		pre.Next = current.Next
	}

	return head
}

// Runtime 2 ms Beats 50.59% Memory 2.2 MB Beats 15.97%
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummpy := &ListNode{Next: head}
	temp := dummpy

	for i := 0; i < n; i++ {
		temp = temp.Next
	}

	pre := dummpy
	for temp.Next != nil {
		temp = temp.Next
		pre = pre.Next
	}
	if pre.Next == head {
		return head.Next
	}
	pre.Next = pre.Next.Next

	return head
}

// 다른사람의 솔루션 메모리 사용을 최소로 했음 사실 아무런 의미가 없는 차이이긴 함
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	first_node := head
	second_node := head
	i := 0
	for first_node != nil {
		if i > n {
			second_node = second_node.Next
		}
		first_node = first_node.Next
		i += 1
	}

	if n == i {
		head = head.Next
	} else {
		second_node.Next = second_node.Next.Next
	}
	return head
}
