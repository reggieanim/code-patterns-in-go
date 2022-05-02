package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution("aabccbb", 2))
}

func Solution(str string, k int) int {
	maxRepeatedLetters := 0
	windowStart := 0
	patternMap := map[string]int{}
	maxLength := 0
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, in := patternMap[string(rightChar)]; !in {
			patternMap[string(rightChar)] = 0
		}
		patternMap[string(rightChar)] += 1
		maxRepeatedLetters = int(math.Max(float64(maxRepeatedLetters), float64(patternMap[string(rightChar)])))
		if windowEnd-windowStart+1-maxRepeatedLetters > k {
			leftChar := str[windowStart]
			patternMap[string(leftChar)] -= 1

			windowStart += 1
		}
		maxLength = int(math.Max(float64(maxLength), float64(windowEnd-windowStart+1)))
	}
	return maxLength
}
