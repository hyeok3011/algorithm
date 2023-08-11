package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			output: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{4, 5},
			},
			output: [][]int{
				{1, 5},
			},
		},
		{
			input: [][]int{
				{1, 4},
				{0, 5},
			},
			output: [][]int{
				{0, 5},
			},
		},
	}

	for _, test := range tests {
		assert.Equal(t, merge(test.input), test.output)
		assert.Equal(t, merge2(test.input), test.output)
	}
}
