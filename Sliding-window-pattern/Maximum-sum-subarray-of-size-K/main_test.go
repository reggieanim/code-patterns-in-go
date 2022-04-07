package main

import "testing"

type cases struct {
	positivIntegers []int
	subArraySize    int
	result          int
}

func TestMaxSumOfSizeK(t *testing.T) {
	tt := []cases{
		{
			[]int{2, 1, 5, 1, 3, 2},
			3,
			9,
		},
		{
			[]int{2, 3, 4, 1, 5},
			2,
			7,
		},
	}
	for _, v := range tt {
		got := MaxSumOfSizeK(v.positivIntegers, v.subArraySize)
		assert(got, v.result, t)
	}
}

func assert(got, result int, t *testing.T) {
	if got == result {
		return
	}
	t.Errorf("got %d, expected %d", got, result)
}
