package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSwapPairs(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	output := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	assert.Equal(t, swapPairs(input), output)
}
