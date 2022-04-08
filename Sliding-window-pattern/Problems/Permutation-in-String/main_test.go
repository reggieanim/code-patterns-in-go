package main

import "testing"

type cases struct {
	str       string
	substring string
	result    bool
}

func TestPermutation(t *testing.T) {
	tt := []cases{
		{
			"oidbcaf",
			"abc",
			true,
		},
		{
			"odicf",
			"dc",
			false,
		},
		{
			"bcdxabcdy",
			"bcdyabcdx",
			true,
		},
		{
			"aaacb",
			"abc",
			true,
		},
	}
	for _, v := range tt {
		got := Permutation(v.str, v.substring)
		assert(got, v.result, v.str, t)
	}
}

func assert(got, result bool, str string, t *testing.T) {
	if got == result {
		return
	}
	t.Errorf("got %v, for '%v' expected %v", got, str, result)
}
