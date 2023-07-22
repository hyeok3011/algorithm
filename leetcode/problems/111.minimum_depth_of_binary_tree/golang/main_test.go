package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinDepth(t *testing.T) {
	testNode := TreeNode{
		Left: &TreeNode{},
		Right: &TreeNode{
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
	}
	assert.Equal(t, minDepth(&testNode), 2)
}
