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

// Runtime 3 ms Beats 90.21% Memory 4.3 MB Beats 23.54%
// preorder : root > left subtree > right subtree
// inorder : left subtree > root > right subtree
// preorder와 inorder의 인덱스 상관관계를 제대로 이해하지 못하고 패턴을 확인해서 풀었다...
// 나중에 다ㅓ시 확인하자 @@@@@@@@
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	currentVal := preorder[0]
	node := &TreeNode{Val: currentVal}

	inorderIndex := 0
	for i, v := range inorder {
		if v == currentVal {
			inorderIndex = i
			break
		}
	}

	node.Left = buildTree(preorder[1:1+inorderIndex], inorder[:inorderIndex])
	node.Right = buildTree(preorder[1+inorderIndex:], inorder[inorderIndex+1:])

	return node
}
