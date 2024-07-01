package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSimplifyPath(t *testing.T) {
	assert.Equal(t, simplifyPath("/home/"), "/home")
	assert.Equal(t, simplifyPath("/../"), "/")
	assert.Equal(t, simplifyPath("/home//foo/"), "/home/foo")
	assert.Equal(t, simplifyPath("/home/../foo/"), "/foo")
	assert.Equal(t, simplifyPath("/a//b////c/d//././/.."), "/a/b/c")

}
