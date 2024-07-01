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

// https://leetcode.com/problems/path-sum-ii/description/
// Runtime 2 ms Beats 93.1% Memory 4.4 MB Beats 97.82%
// 메모리를 최대한 조금 먹으려고 별짓을 다해봤지만 100퍼는 나오지 않음.... 뭔가 개선이 필요함
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	dfs(root, &[]int{targetSum}, &result)
	return result
}

func dfs(node *TreeNode, stack *[]int, result *[][]int) {
	if node == nil {
		return
	}

	(*stack)[0] -= node.Val
	*stack = append(*stack, node.Val)
	if (*stack)[0] == 0 && node.Left == nil && node.Right == nil {
		(*stack)[0] += node.Val
		*result = append(*result, append([]int{}, (*stack)[1:]...))
		*stack = (*stack)[:len(*stack)-1]
		return
	}

	if node.Left != nil {
		dfs(node.Left, stack, result)
	}

	if node.Right != nil {
		dfs(node.Right, stack, result)
	}

	*stack = (*stack)[:len(*stack)-1]
	(*stack)[0] += node.Val
	return
}

func pathSumDFSStack(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	answer := [][]int{}

	nodeStack := []*TreeNode{root}
	sumPathsStack := [][]int{{0}}

	for len(nodeStack) != 0 {
		currentNode := nodeStack[len(nodeStack)-1]
		sumPath := sumPathsStack[len(sumPathsStack)-1]

		nodeStack = nodeStack[:len(nodeStack)-1]
		sumPathsStack = sumPathsStack[:len(sumPathsStack)-1]

		sumPath[0] += currentNode.Val
		sumPath = append(sumPath, currentNode.Val)
		if sumPath[0] == targetSum && currentNode.Left == nil && currentNode.Right == nil {
			answer = append(answer, sumPath[1:])
		}

		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
			sumPathsStack = append(sumPathsStack, sumPath)
		}

		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
			sumPathsStack = append(sumPathsStack, sumPath)
		}
	}

	return answer
}
