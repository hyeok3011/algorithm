package main

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
		{
			str1: "ABCDEF",
			str2: "ABC",
			want: "",
		},
		{
			str1: "TAUXXTAUXXTAUXXTAUXXTAUXX",
			str2: "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			want: "TAUXX",
		},
		{
			str1: "ABABABAB",
			str2: "ABAB",
			want: "ABAB",
		},
		{
			str1: "AAAAAAAAA",
			str2: "AAACCC",
			want: "",
		},
	}

	for _, test := range tests {
		result := gcdOfStrings(test.str1, test.str2)
		if result != test.want {
			t.Errorf("gcdOfStrings(%s, %s) = %s; want %s", test.str1, test.str2, result, test.want)
		}
	}

}
