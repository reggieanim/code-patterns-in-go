package main

import (
	"reflect"
	"testing"
)

type cases struct {
	str       string
	substring string
	result    []int
}

func TestAnagrams(t *testing.T) {
	tt := []cases{
		{
			"ppqp",
			"pq",
			[]int{1, 2},
		},
		{
			"abbcabc",
			"abc",
			[]int{2, 3, 4},
		},
	}
	for _, v := range tt {
		got := Anagrams(v.str, v.substring)
		assert(got, v.result, v.str, t)
	}
}

func assert(got, result []int, str string, t *testing.T) {
	if reflect.DeepEqual(got, result) {
		return
	}
	t.Errorf("got %v, for '%v' expected %v", got, str, result)
}
