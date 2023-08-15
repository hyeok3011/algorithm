package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/binary-tree-postorder-traversal/description/
// Runtime 1 ms Beats 72.35% Memory 2 MB Beats 100%
// 방법마다 속도와 메모리는 큰 차이가 없음 유동적으로 왔다갔자 하는 상태 0ms ~ 1ms
// 제일 효율적이라고 생각되는 방법은 Recursion이라고 생각중.

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	values := []int{}
	nodeStack := []*TreeNode{root}

	for len(nodeStack) > 0 {
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		values = append([]int{currentNode.Val}, values...)
		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
		}

		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
		}
	}

	return values
}
func postorderTraversalRecursion(root *TreeNode) []int {
	values := []int{}

	traversal(root, &values)

	return values
}

func traversal(node *TreeNode, values *[]int) {
	if node == nil {
		return
	}

	traversal(node.Left, values)
	traversal(node.Right, values)
	*(values) = append(*(values), node.Val)

}

func postorderTraversalStack2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	values := []int{}
	nodeStack := []*TreeNode{root}

	for len(nodeStack) > 0 {
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		values = append(values, currentNode.Val)

		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
		}

		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
		}
	}

	reverse(values)

	return values
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
