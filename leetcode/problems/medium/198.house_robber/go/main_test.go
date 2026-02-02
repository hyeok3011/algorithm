package main

import (
	"testing"
)

func TestRob(t *testing.T) {
	testCases := []struct {
		Input    []int
		Out      int
		ErrorMsg string
	}{
		{Input: []int{1, 2, 3, 1}, Out: 4, ErrorMsg: ""},
		{Input: []int{2, 7, 9, 3, 1}, Out: 12, ErrorMsg: ""},
	}

	var result int
	for _, tc := range testCases {
		result = rob(tc.Input)
		if tc.Out != result {
			t.Errorf("expected: |%v got: |%v error: %s", tc.Out, result, tc.ErrorMsg)
		}
	}

}
