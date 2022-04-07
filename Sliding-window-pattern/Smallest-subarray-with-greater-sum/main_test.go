package main

import "testing"

type cases struct {
	positivIntegers []int
	subArraySize    int
	result          int
}

func TestSmallestSubArrWGreaterSum(t *testing.T) {
	tt := []cases{
		{
			[]int{2, 1, 5, 2, 3, 2},
			7,
			2,
		},
		{
			[]int{2, 1, 5, 2, 8},
			7,
			1,
		},
		{
			[]int{2, 1, 5, 2, 8},
			500,
			0,
		},
	}
	for _, v := range tt {
		got := SmallestSubArrWGreaterSum(v.positivIntegers, v.subArraySize)
		assert(got, v.result, t)
	}
}

func assert(got, result int, t *testing.T) {
	if got == result {
		return
	}
	t.Errorf("got %d, expected %d", got, result)
}
