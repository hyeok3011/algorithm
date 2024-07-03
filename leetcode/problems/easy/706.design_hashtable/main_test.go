package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMyHashMap(t *testing.T) {
	hashMap := Constructor()
	// ["MyHashMap","put","put","get","get","put","get","remove","get"]
	// [[],[1,1],[2,2],[1],[3],[2,1],[2],[2],[2]]
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	assert.Equal(t, hashMap.Get(1), 1)
	assert.Equal(t, hashMap.Get(3), -1)
	hashMap.Put(2, 1)
	assert.Equal(t, hashMap.Get(2), 1)
	hashMap.Remove(2)
	assert.Equal(t, hashMap.Get(2), -1)
}
