package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	prevNode := &TreeNode{}

	var recursion func(node *TreeNode)
	recursion = func(node *TreeNode) {
		if node == nil {
			return
		}
		prevNode.Right = node
		temp := node.Right
		prevNode = node
		recursion(node.Left)
		recursion(temp)
		node.Left = nil
	}

	recursion(root)
}

func flattenStack(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	var prevNode *TreeNode

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prevNode != nil {
			prevNode.Right = node
			prevNode.Left = nil
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		prevNode = node
	}
}
