package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
	}

	for _, c := range cases {
		moveZeroes2(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("moveZeroes(%v) = %v, expected %v", c.input, c.input, c.expected)
		}
	}
}
