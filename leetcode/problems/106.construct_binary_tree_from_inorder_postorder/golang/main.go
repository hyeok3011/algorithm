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

// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
// 비슷한 문제를 어제 풀었지만 풀때마다 헷갈리는 문제@@@
// Runtime 6 ms Beats 53.76% Memory 4 MB Beats 73.41%

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	currentVal := postorder[len(postorder)-1]
	node := &TreeNode{Val: currentVal}

	inorderIndex := 0
	for i, v := range inorder {
		if v == currentVal {
			inorderIndex = i
			break
		}
	}

	node.Left = buildTree(inorder[:inorderIndex], postorder[:inorderIndex])
	node.Right = buildTree(inorder[inorderIndex+1:], postorder[inorderIndex:len(postorder)-1])
	return node
}

// inorder의 값의 인덱스를 for문으로 찾지않고 미리 map으로 만들어두는 솔루션도 있음
func AnotherSolution(inorder []int, postorder []int) *TreeNode {
	inorderIndexMap := make(map[int]int)

	for k, v := range inorder {
		inorderIndexMap[v] = k
	}

	return toTree(postorder, inorder, 0, len(inorder)-1, inorderIndexMap)
}

func toTree(po, io []int, start, end int, iim map[int]int) *TreeNode {
	if len(po) == 0 {
		return nil
	}

	rootVal := po[len(po)-1]
	rootIndex := iim[rootVal]
	leftNodes := rootIndex - start

	root := &TreeNode{rootVal, nil, nil}
	root.Left = toTree(po[:leftNodes], io, start, rootIndex-1, iim)
	root.Right = toTree(po[leftNodes:len(po)-1], io, rootIndex+1, end, iim)

	return root
}
