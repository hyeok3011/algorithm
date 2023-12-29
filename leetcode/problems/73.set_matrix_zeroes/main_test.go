package main

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
			},
			output: [][]int{
				{1, 0, 1}, {0, 0, 0}, {1, 0, 1},
			},
		},
		{
			input: [][]int{
				{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
			},
			output: [][]int{
				{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0},
			},
		},
	}

	for _, tc := range testCases {
		setZeroes(tc.input)
		if !reflect.DeepEqual(tc.input, tc.output) {
			t.Fatalf("expected: %v, got: %v", tc.output, tc.input)
		}
	}

}
