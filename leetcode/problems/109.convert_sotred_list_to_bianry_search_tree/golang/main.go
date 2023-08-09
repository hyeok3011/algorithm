package main

// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/description/
// Runtime 4 ms Beats 82.81% Memory 5.9 MB Beats 92.19%
// O(nlogn)
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	midListNode := findMiddleNodeAndDisconnection(head)
	currentTreeNode := &TreeNode{Val: midListNode.Val}

	if midListNode == head {
		return currentTreeNode
	}

	currentTreeNode.Left = sortedListToBST(head)
	currentTreeNode.Right = sortedListToBST(midListNode.Next)

	return currentTreeNode
}

func findMiddleNodeAndDisconnection(headNode *ListNode) *ListNode {
	slow, fast, pre := headNode, headNode, headNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil

	return slow
}

// 흠.... 시간 복잡도를 줄여보려고 미리 slice에 담아봤는데 큰 의미가 없다. 추출하는 과정에서 이미
// O(n)을 추가로 잡아먹어서 그런듯 하다..
func sortedListToBST2(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return buildBST(0, len(nums)-1, nums)
}

func buildBST(startIndex, endIndex int, nums []int) *TreeNode {
	if startIndex > endIndex {
		return nil
	}

	midIndex := startIndex + (endIndex-startIndex)/2

	currentTreeNode := &TreeNode{Val: nums[midIndex]}

	currentTreeNode.Left = buildBST(startIndex, midIndex-1, nums)
	currentTreeNode.Right = buildBST(midIndex+1, endIndex, nums)

	return currentTreeNode
}
