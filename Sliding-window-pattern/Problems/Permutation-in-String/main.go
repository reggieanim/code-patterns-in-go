package main

import "fmt"

func main() {}

func Permutation(str, substr string) bool {
	matched := 0
	substrMap := map[string]int{}
	windowStart := 0

	for _, v := range substr {
		if _, in := substrMap[string(v)]; !in {
			substrMap[string(v)] = 0
		}
		substrMap[string(v)] += 1
	}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		fmt.Println(substrMap)
		rightChar := string(str[windowEnd])
		if _, in := substrMap[rightChar]; in {
			substrMap[rightChar] -= 1
			if substrMap[rightChar] == 0 {
				matched += 1
			}
		}

		if matched == len(substrMap) {
			return true
		}

		if windowEnd >= len(substr)-1 {
			leftChar := string(str[windowStart])
			windowStart += 1
			if _, in := substrMap[leftChar]; in {
				if substrMap[leftChar] == 0 {
					matched -= 1
				}
				substrMap[leftChar] += 1
			}
		}

	}
	return false
}
