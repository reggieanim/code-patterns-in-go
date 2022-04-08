package main

import "testing"

type cases struct {
	word   string
	k      int
	result int
}

func TestLongestSubStrDis(t *testing.T) {
	tt := []cases{
		{
			"araaci",
			2,
			4,
		},
		{
			"araaci",
			1,
			2,
		},
		{
			"cbbebi",
			3,
			5,
		},
		{
			"cbbebi",
			10,
			6,
		},
	}
	for _, v := range tt {
		got := LongestSubStrDis(v.word, v.k)
		assert(got, v.result, t)
	}
}

func assert(got, result int, t *testing.T) {
	if got == result {
		return
	}
	t.Errorf("got %d, expected %d", got, result)
}
