package main

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/description/
func pairSum(head *ListNode) int {
	slow, fast := head, head
	halfLen := 0
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		halfLen++
		fast = fast.Next.Next
	}

	middleNode := reverseNode(slow)
	maxSum := 0
	for i := 0; i < halfLen; i++ {
		maxSum = max(maxSum, head.Val + middleNode.Val)
		head = head.Next
		middleNode = middleNode.Next
	}
	return maxSum
}

func reverseNode(node *ListNode) *ListNode {
	next := node.Next
	for next != nil {
		temp := next.Next
		next.Next = node
		node = next
		next = temp
	}
	
	return node
}
