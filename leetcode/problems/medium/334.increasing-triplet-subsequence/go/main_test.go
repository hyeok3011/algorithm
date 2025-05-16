package main

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		// {[]int{1, 2, 3, 4, 5}, true},
		{[]int{2, 1, 5, 0, 4, 6}, true},
	}

	for _, test := range tests {
		if got := increasingTriplet(test.nums); got != test.want {
			t.Errorf("increasingTriplet(%v) = %v, want %v", test.nums, got, test.want)
		}
	}
}
