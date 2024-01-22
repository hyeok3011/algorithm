package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	assert.Equal(t, longestValidParentheses("(()"), 2)
	assert.Equal(t, longestValidParentheses("()"), 2)
	assert.Equal(t, longestValidParentheses("(()()"), 4)
	assert.Equal(t, longestValidParentheses("(()())()"), 8)
	assert.Equal(t, longestValidParentheses("()(()"), 2)
	assert.Equal(t, longestValidParentheses(")(((((()())()()))()(()))("), 22)
	assert.Equal(t, longestValidParentheses(")()())"), 4)
	assert.Equal(t, longestValidParentheses(")()"), 2)

}
