package main

import (
	"fmt"
	"math"
)

func main() {
	result := Solution([]int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}, 2)
	fmt.Println(result)
}

func Solution(arr []int, k int) int {
	windowStart := 0
	maximumLength := 0
	repeatedChars := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		rightChar := arr[windowEnd]
		if rightChar == 1 {
			repeatedChars += 1
		}
		if windowEnd-windowStart+1-repeatedChars > k {
			leftChar := arr[windowStart]
			if leftChar == 1 {
				repeatedChars -= 1
			}
			windowStart += 1
		}
		maximumLength = int(math.Max(float64(windowEnd-windowStart+1), float64(maximumLength)))
	}
	return maximumLength
}
