package main

import "testing"

func createTreeFromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		tree1    []int
		tree2    []int
		expected bool
	}{
		{
			tree1:    []int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4},
			tree2:    []int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8},
			expected: true,
		},
		{
			tree1:    []int{1, 2, 3},
			tree2:    []int{1, 3, 2},
			expected: false,
		},
		{
			tree1:    []int{1},
			tree2:    []int{1},
			expected: true,
		},
		{
			tree1:    []int{},
			tree2:    []int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		tree1 := createTreeFromArray(tt.tree1)
		tree2 := createTreeFromArray(tt.tree2)
		result := leafSimilar(tree1, tree2)

		if result != tt.expected {
			t.Errorf("got %v, want %v", result, tt.expected)
		}
	}
}
