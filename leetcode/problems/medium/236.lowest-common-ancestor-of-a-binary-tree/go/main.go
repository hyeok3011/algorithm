package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

import (
	"github.com/hyeok3011/algorithm/structures"
)

// TreeNode는 structures 패키지의 것을 사용한다고 가정합니다.
type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var findPath func(node *TreeNode, target *TreeNode, path *[]*TreeNode) bool
	findPath = func(node *TreeNode, target *TreeNode, path *[]*TreeNode) bool {
		if node == nil {
			return false
		}

		*path = append(*path, node)

		if node == target {
			return true
		}

		if findPath(node.Left, target, path) || findPath(node.Right, target, path) {
			return true
		}

		*path = (*path)[:len(*path)-1]
		return false
	}

	pathToP := []*TreeNode{}
	findPath(root, p, &pathToP)

	pathToQ := []*TreeNode{}
	findPath(root, q, &pathToQ)

	var lca *TreeNode
	i := 0

	for i < len(pathToP) && i < len(pathToQ) && pathToP[i] == pathToQ[i] {
		lca = pathToP[i]
		i++
	}

	return lca
}

// 다른사람의 풀이......
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
