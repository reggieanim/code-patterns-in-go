package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution("aabdec", "abc"))
}

func Solution(str, pattern string) string {
	patternMap := map[string]int{}
	for _, v := range pattern {
		if _, in := patternMap[string(v)]; !in {
			patternMap[string(v)] = 0
		}
		patternMap[string(v)] += 1
	}
	windowStart := 0
	match := 0
	currentLength := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]
		if _, in := patternMap[string(rightChar)]; in {
			patternMap[string(rightChar)] -= 1
			if patternMap[string(rightChar)] == 0 {
				match += 1
			}
		}

		for match == len(pattern) {
			leftChar := str[windowStart]
			if _, in := patternMap[string(leftChar)]; in {
				patternMap[string(leftChar)] += 1
				if patternMap[string(leftChar)] == 0 {
					match += -1
				}
			}
			fmt.Println("start", windowEnd)
			fmt.Println("end", windowStart)
			currentLength = int(math.Max(float64(currentLength), float64(windowEnd-windowStart+1)))
			windowStart += 1
		}
	}
	return str[windowStart:currentLength]
}
