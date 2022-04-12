package main

import "fmt"

func main() {

}

func Anagrams(str, pattern string) []int {
	anagramIndices := []int{}
	windowStart := 0
	match := 0
	patternMap := map[string]int{}

	for _, v := range pattern {
		patternMap[string(v)] += 1
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChar := str[windowEnd]

		if _, in := patternMap[string(rightChar)]; in {

			patternMap[string(rightChar)] -= 1
			if patternMap[string(rightChar)] == 0 {
				match += 1
			}

			fmt.Println(patternMap)
		}

		if match == len(pattern) {
			anagramIndices = append(anagramIndices, windowStart)
		}

		if windowEnd+1 >= len(pattern) {
			//remove from windowStart
			leftChar := str[windowStart]
			if _, in := patternMap[string(leftChar)]; in {
				if patternMap[string(leftChar)] == 0 {
					match -= 1
				}
				patternMap[string(leftChar)] += 1
			}
			windowStart += 1
		}
	}

	return anagramIndices

}
