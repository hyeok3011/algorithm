package main

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := subsetsWithDup(test.nums)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("subsetsWithDup(%v) = %v, want %v", test.nums, got, test.want)
			}
		})
	}
}
