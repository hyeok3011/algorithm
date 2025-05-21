package main

import (
	"reflect"
	"testing"
)

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

func TestRightSideView(t *testing.T) {
	tests := []struct {
		tree     []int
		expected []int
	}{
		{
			tree:     []int{1, 2, 3, -1, 5, -1, 4},
			expected: []int{1, 3, 4},
		},
		{
			tree:     []int{1, -1, 3},
			expected: []int{1, 3},
		},
		{
			tree:     []int{},
			expected: []int{},
		},
		{
			tree:     []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		tree := createTreeFromArray(tt.tree)
		result := rightSideView(tree)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("got %v, want %v", result, tt.expected)
		}
	}
}
