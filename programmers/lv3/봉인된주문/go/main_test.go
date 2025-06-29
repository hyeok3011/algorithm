package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		n    int64
		bans []string
		want string
	}{
		// {30, []string{"d", "e", "bb", "aa", "ae"}, "ah"},
		// {7388, []string{"gqk", "kdn", "jxj", "jxi", "fug", "jxg", "ewq", "len", "bhc"}, "jxk"},
		{52, []string{}, "az"},
	}
	for _, test := range tests {
		got := solution(test.n, test.bans)
		if got != test.want {
			t.Errorf("solution(%d, %v) = %s, want %s", test.n, test.bans, got, test.want)
		}
	}
}
