package main

import (
	"math"
)

func main() {

}

func LongestSubStrDis(str string, k int) int {
	windowStart := 0
	maximumLength := 0
	distinctChars := map[string]int{}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := string(str[windowEnd])
		if _, in := distinctChars[rightChar]; !in {
			distinctChars[rightChar] = 0
		}
		distinctChars[rightChar] += 1

		for len(distinctChars) > k {
			leftChar := string(str[windowStart])
			distinctChars[leftChar] -= 1
			if distinctChars[leftChar] == 0 {
				delete(distinctChars, leftChar)
			}
			windowStart += 1
		}

		maximumLength = int(math.Max(float64(maximumLength), float64(windowEnd-windowStart+1)))
	}

	return maximumLength
}
