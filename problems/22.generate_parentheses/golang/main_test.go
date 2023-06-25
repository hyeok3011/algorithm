package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGenerateParenthesis(t *testing.T) {
	assert.Equal(t, generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	assert.Equal(t, generateParenthesis(1), []string{"()"})
}
