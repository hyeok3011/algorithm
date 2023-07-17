package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
// Runtime 0 ms Beats 100% Memory 3.1 MB Beats 25.84%
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return head
}
